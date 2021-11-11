package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

// todo move to config
var (
	authServerAddr = flag.String("auth_server_addr", "localhost:8878", "The server address in the format of host:port")
	taskServerAddr = flag.String("task_server_addr", "localhost:8879", "The server address in the format of host:port")
)

type gateway struct {
	authClient auth.AuthClient
}

func SessionMiddleware(authClient auth.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		//token, err := c.Cookie("session_token")
		//if err != nil {
		//
		//}
		//
		//// todo
		//ctx := context.Background()
		//authClient.CheckSession(ctx, &auth.CheckSessionRequest{
		//
		//})

		//fmt.Println(token) // todo check session token
	}
}

func main() {
	r := gin.Default()
	//var opts []grpc.DialOption
	conn, err := grpc.Dial(*authServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create auth server conn: %v", err)
	}

	gateway := &gateway{
		authClient: auth.NewAuthClient(conn),
	}

	// create auth service client
	// create tasks service client
	r.Use(SessionMiddleware(gateway.authClient))
	// auth
	r.POST("/api/v1/login", func(c *gin.Context) {
		ctx := context.Background()
		req := &struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(req); err != nil {
			// todo
		}

		resp, err := gateway.authClient.Login(ctx, &auth.LoginRequest{
			Login:    req.Login,
			Password: req.Password,
		})

		if err != nil {
			log.Fatalf("failed to create ")
		}

		c.SetCookie("session_token", resp.Token, 60*60*24, "/", "localhost", false, true)
	})

	type userInfoJSON struct {
		Login    *string `json:"login"`
		Password *string `json:"password"`
		Email    *string `json:"email"`
		FullName *string `json:"full_name"`
		Role     *string `json:"role"`
	}

	r.POST("/api/v1/register", func(c *gin.Context) {
		req := userInfoJSON{}

		if err := c.ShouldBind(&req); err != nil {
			errr := err.Error()
			fmt.Println(errr)
			c.Status(500)
			return
		}

		ctx := context.Background()
		resp, err := gateway.authClient.Register(ctx, &auth.RegisterRequest{UserInfo: &auth.UserInfo{
			Login:    &auth.StringContainer{Value: *req.Login},
			Password: &auth.StringContainer{Value: *req.Password},
			Email:    &auth.StringContainer{Value: *req.Email},
			FullName: &auth.StringContainer{Value: *req.FullName},
			Role:     &auth.StringContainer{Value: *req.Role},
		}})
		if err != nil {
			// handle err
		}

		fmt.Println(resp)
		c.String(http.StatusOK, "ok")
	})

	// users
	r.GET("/api/v1/user/get/:id", func(c *gin.Context) {
		//gateway.authClient.GetUserById()
	})

	r.PUT("/api/v1/user/update/:id", func(c *gin.Context) {
		//gateway.authClient.UpdateUserById()
	})
	r.DELETE("/api/v1/user/delete/:id", func(c *gin.Context) {
		//gateway.authClient.DeleteUserById()
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

	r.Run()
}
