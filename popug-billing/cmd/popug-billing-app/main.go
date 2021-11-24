package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/popug-billing/internal/api"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
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

	postgresDsn := "host=localhost user=postgres password=password dbname=tasks port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	impl := api.New()

	var billingCloseCh chan struct{}

	billingCloseCh = startBillingCycle()
	// todo map of channels is needed
	msgCh := make(chan *kafka.Message)

	go func() {
		select {
		case msg := <-msgCh:
			switch *msg.TopicPartition.Topic {
			case completeTaskTopic:
				impl.CreatePlusTransaction()
			case assignTaskTopic:
				impl.CreateMinusTransaction()
			case "task-create-events":
				// on create event we generate prices
				// what if on create transaction we don't have task exist?
				l := -20
				u := -10
				cost := l + rand.Intn(u-l+1)

				l = 20
				u = 40
				prize := l + rand.Intn(u-l+1)
				impl.CreateTask(ctx, cost, prize)
			}
		case <-billingCloseCh:
			impl.CloseBillingCycle()
		}
	}()

	go subscribe(msgCh)

	// todo close billing cycle on ticker

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
		//"task-create-events",
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
