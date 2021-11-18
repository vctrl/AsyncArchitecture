package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/api"
	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/db"
	"github.com/vctrl/async-architecture/week_2/schema/events"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	postgresDsn := "host=localhost user=postgres password=password dbname=tasks port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	impl := api.New(postgresDsn)
	tasks.RegisterTasksServer(srv, impl)

	msgCh := make(chan *kafka.Message)

	// handling events
	go func() {
		users := db.NewUserRepositorySQL(postgresDsn)
		for msg := range msgCh {
			context := context.Background()
			switch *msg.TopicPartition.Topic {
			// todo move to consts
			case "user-create-events":
				evt := events.UserCreatedEvent{}
				err = proto.Unmarshal(msg.Value, &evt)

				err := users.Create(context, &db.User{
					PublicID: evt.GetPublicId(),
					Login:    evt.GetLogin(),
					Email:    evt.GetEmail(),
					FullName: evt.GetFullName(),
				})

				if err != nil {
					fmt.Printf("error create user record: %v\n", err.Error())
					// todo handle err
				}
			case "user-update-events":
				// todo update record
			case "user-delete-events":
				// todo delete record
			default:
				// todo handle not supported event type
			}
		}
	}()

	// todo: the total mess is coming
	// todo interface: hide implementation details
	// todo create consumer struct
	// todo handle error
	go subscribe(msgCh)

	if err != nil {
		panic(err)
	}

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
