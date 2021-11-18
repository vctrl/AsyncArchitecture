package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/api"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8878))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	kafkaCfg := &kafka.ConfigMap{"bootstrap.servers": "localhost:29092"}
	postgresDsn := "host=localhost user=postgres password=password dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	impl, err := api.New(postgresDsn, kafkaCfg)
	if err != nil {
		log.Fatalf("failed to create server impl: %s", err)
	}

	auth.RegisterAuthServer(srv, impl)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
