package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	"github.com/vctrl/async-architecture/popug-billing/internal/api"
	"github.com/vctrl/async-architecture/popug-billing/internal/db"
	"github.com/vctrl/async-architecture/schema/billing"
	"github.com/vctrl/async-architecture/schema/events"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	completeTaskTopic = "task-complete-events"
	assignTaskTopic   = "task-assign-events"
)

func main() {
	lis, err := net.Listen("tcp", ":8881")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	postgresDsn := "host=localhost user=postgres password=password dbname=billing port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	impl := api.New(postgresDsn)

	billing.RegisterBillingServer(srv, impl)
	ticker := time.NewTicker(time.Minute)

	// todo map of channels is needed
	msgCh := make(chan *kafka.Message)

	tasks := db.NewTaskRepoSQL(postgresDsn)

	go func() {
		ctx := context.Background()
		for {
			// todo cancel with context
			select {
			case <-ticker.C:
				impl.Mdl.CloseBillingCycle(ctx)
			case msg := <-msgCh:
				switch *msg.TopicPartition.Topic {
				case completeTaskTopic:
					//todo logging
					fmt.Println("read complete task event")
					evt := events.TaskCompletedEvent{}
					err = proto.Unmarshal(msg.Value, &evt)
					if err != nil {
						panic(err)
					}
					// create transaction, bonus from task
					_, _, err = impl.Mdl.CreatePlusTransaction(ctx, evt.TaskPublicId, evt.AssignedToPublicId)
					if err != nil {
						panic(err)
					}
				case assignTaskTopic:
					fmt.Println("read assign task event")
					evt := events.TaskAssignedEvent{}
					err = proto.Unmarshal(msg.Value, &evt)

					if err != nil {
						panic(err)
					}
					// create transaction, take cost from task
					_, _, err := impl.Mdl.CreateMinusTransaction(ctx, evt.TaskPublicId, evt.AssignedToPublicId)
					if err != nil {
						panic(err)
					}
				case "task-create-events":
					// on create event we generate prices
					// what if on create transaction we don't have task exist?
					fmt.Println("read create task event")
					l := -20
					u := -10
					cost := l + rand.Intn(u-l+1)

					l = 20
					u = 40
					prize := l + rand.Intn(u-l+1)
					evt := events.TaskCreatedEvent{}
					err = proto.Unmarshal(msg.Value, &evt)

					if err != nil {
						panic(err)
					}
					_, _, err = tasks.Create(ctx, &db.Task{
						ID:       uuid.NewV4().String(),
						PublicID: evt.PublicId,
						Prize:    prize,
						Cost:     cost,
					})

					if err != nil {
						fmt.Printf("error create user record: %v\n", err.Error())
						// todo handle err
					}
				}
			}
		}

	}()

	go subscribe(msgCh)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}

func subscribe(msgCh chan *kafka.Message) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "myGroup228",
		"auto.offset.reset": "beginning",
	})

	if err != nil {
		return err
	}

	err = c.SubscribeTopics([]string{
		"user-create-events",
		//"user-update-events",
		//"user-delete-events",
		"task-create-events",
		//"task-update-events",
		//"task-delete-events",
		completeTaskTopic,
		assignTaskTopic,
	}, nil)

	if err != nil {
		return err
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			msgCh <- msg
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	return nil
}
