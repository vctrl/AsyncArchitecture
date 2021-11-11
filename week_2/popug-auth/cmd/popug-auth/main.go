package main

import (
	"fmt"
	"github.com/vctrl/async-architecture/week_2/popug-auth/internal/api"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

// todo move to config
//var (
//	authServerAddr = flag.String("server_addr", "localhost:8878", "The server address in the format of host:port")
//	taskServerAddr = flag.String("server_addr", "localhost:8879", "The server address in the format of host:port")
//)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8878))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	impl, err := api.New()
	if err != nil {
		log.Fatalf("failed to create server impl: %s", err)
	}

	auth.RegisterAuthServer(srv, impl)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
