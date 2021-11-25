package main

import (
	"github.com/vctrl/async-architecture/schema/analytics"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	postgresDsn := "host=localhost user=postgres password=password dbname=tasks port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	impl := api.New(postgresDsn)
	analytics.Register(srv, impl)
}
