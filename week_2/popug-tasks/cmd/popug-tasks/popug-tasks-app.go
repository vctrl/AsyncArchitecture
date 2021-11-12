package main

import (
	"github.com/vctrl/async-architecture/week_2/popug-tasks/internal/api"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
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
	impl := api.NewServer()
	tasks.RegisterTasksServer(srv, impl)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
