package main

import (
	"log"
	"net"

	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/api"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	impl := api.New("host=localhost user=postgres password=password dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	tasks.RegisterTasksServer(srv, impl)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
