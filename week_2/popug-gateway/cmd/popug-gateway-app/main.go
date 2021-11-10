package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc"
)
var (
	serverAddr = flag.String("server_addr", "localhost:8878", "The server address in the format of host:port")
)

type gateway struct {
	// auth
	// tasks
	// logger
}

func SessionMiddleware(authClient auth.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session_token")
		if err != nil {

		}

		// todo
		//authClient.Check(token)

		fmt.Println(token) // todo check session token
	}
}

func main() {
	r := gin.Default()
	var opts []grpc.DialOption
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {

	}
	authClient := auth.NewAuthClient(conn)
	// create auth service client
	// create tasks service client
	r.Use(SessionMiddleware(authClient))

	// auth
	r.POST("/api/v1/login", func(c *gin.Context) {
		ctx := context.Background()
		authClient.Login(ctx, &auth.LoginRequest{
			// todo
			Login:    "",
			Password: "",
		})
		c.SetCookie("session_token", "token", 60*60*24, "/", "localhost", false, true)
	})
	r.POST("/api/v1/register", func(c *gin.Context) {
		// AuthServiceClient.Register()
	})

	// users
	r.GET("/api/v1/user/get/:id", func(c *gin.Context) {

	})
	r.POST("/api/v1/user/create", func(c *gin.Context) {

	})
	r.PUT("/api/v1/user/update/:id", func(c *gin.Context) {

	})
	r.DELETE("/api/v1/user/delete/:id", func(c *gin.Context) {

	})

	//tasks
	r.GET("/api/v1/tasks", func(c *gin.Context) {

	})
	r.GET("/api/v1/task/:id", func(c *gin.Context) {

	})
	r.POST("/api/v1/task/create", func(c *gin.Context) {

	})
	r.PUT("/api/v1/task/update/:id", func(c *gin.Context) {

	})
	r.DELETE("/api/v1/task/delete/:id", func(c *gin.Context) {

	})
	r.POST("/api/v1/task/assign/:id", func(c *gin.Context) {

	})
	r.POST("/api/v1/tasks/shuffle", func(c *gin.Context) {

	})
}
