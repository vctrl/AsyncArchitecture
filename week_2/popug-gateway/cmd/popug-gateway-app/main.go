package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/vctrl/async-architecture/week_2/schema/auth"
	"github.com/vctrl/async-architecture/week_2/schema/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net/http"
	"strings"
	"time"
)

// todo move to config
var (
	authServerAddr = flag.String("auth_server_addr", "localhost:8878", "The server address in the format of host:port")
	taskServerAddr = flag.String("task_server_addr", "localhost:8879", "The server address in the format of host:port")
)

type gateway struct {
	authClient  auth.AuthClient
	tasksClient tasks.TasksClient
}

func SessionMiddleware(authClient auth.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch {
		case strings.HasSuffix(c.Request.URL.Path, "/login"):
			c.Next()
			return
		case strings.HasSuffix(c.Request.URL.Path, "/register"):
			c.Next()
			return
		}

		token, err := c.Cookie("session_token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		resp, err := authClient.CheckSession(ctx, &auth.CheckSessionRequest{
			Token: token,
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if resp.Status.GetCode() != int32(codes.OK) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	aConn, err := grpc.Dial(*authServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create auth server conn: %v", err)
	}
	tConn, err := grpc.Dial(*taskServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create tasks server conn: %v", err)
	}

	gw := &gateway{
		authClient:  auth.NewAuthClient(aConn),
		tasksClient: tasks.NewTasksClient(tConn),
	}

	r.Use(SessionMiddleware(gw.authClient))

	// auth
	r.POST("/api/v1/login", func(c *gin.Context) {
		ctx := context.Background()
		req := &struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		resp, err := gw.authClient.Login(ctx, &auth.LoginRequest{
			Login:    req.Login,
			Password: req.Password,
		})

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.SetCookie("session_token", resp.Token, 60*60*24, "/", "localhost", false, true)
	})

	type user struct {
		Login    *string `json:"login"`
		Password *string `json:"password"`
		Email    *string `json:"email"`
		FullName *string `json:"full_name"`
		Role     *string `json:"role"`
	}

	r.POST("/api/v1/register", func(c *gin.Context) {
		req := user{}

		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx := context.Background()
		resp, err := gw.authClient.Register(ctx, &auth.RegisterRequest{UserInfo: &auth.UserInfo{
			Login:    &auth.StringContainer{Value: *req.Login},
			Password: &auth.StringContainer{Value: *req.Password},
			Email:    &auth.StringContainer{Value: *req.Email},
			FullName: &auth.StringContainer{Value: *req.FullName},
			Role:     &auth.StringContainer{Value: *req.Role},
		}})

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, resp.Id)
	})

	// users
	r.GET("/api/v1/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := gw.authClient.GetUserById(ctx, &auth.GetUserByIdRequest{Id: id})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, resp.UserInfo)
		return
	})

	r.PATCH("/api/v1/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := user{}

		if err := c.ShouldBind(&req); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		gw.authClient.UpdateUserById(ctx, &auth.UpdateUserByIdRequest{
			Id: id,
			UserInfo: &auth.UserInfo{
				Login:    &auth.StringContainer{Value: *req.Login},
				Password: &auth.StringContainer{Value: *req.Password},
				Email:    &auth.StringContainer{Value: *req.Email},
				FullName: &auth.StringContainer{Value: *req.FullName},
				Role:     &auth.StringContainer{Value: *req.Role},
			},
		})
	})

	r.DELETE("/api/v1/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := gw.authClient.DeleteUserById(ctx, &auth.DeleteUserByIdRequest{Id: id})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if resp.Status.GetCode() != int32(codes.OK) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
		return
	})

	//tasks
	r.POST("/api/v1/task", func(c *gin.Context) {
		type task struct {
			Description  string `json:"description"`
			AssignedToID string `json:"assigned_to_id"`
		}

		req := &task{}
		if err := c.ShouldBind(req); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := gw.tasksClient.CreateAndAssignTo(ctx, &tasks.CreateAndAssignToRequest{
			TaskInfo: &tasks.TaskInfo{
				Description: req.Description,
				Done:        false,
			},
			AssignToId: req.AssignedToID,
		})

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if resp.Status.GetCode() != int32(codes.OK) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, &struct{ id string }{resp.Id})
	})

	r.POST("/api/v1/tasks/shuffle", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := gw.tasksClient.Shuffle(ctx, &tasks.ShuffleRequest{})
		// todo move error handling to func
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if resp.Status.GetCode() != int32(http.StatusOK) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
		return
	})

	r.POST("/api/v1/task/:id/done", func(c *gin.Context) {
		id := c.Param("id")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := gw.tasksClient.MarkAsDone(ctx, &tasks.MarkAsDoneRequest{TaskId: id})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if resp.Status.GetCode() != int32(codes.OK) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
		return
	})

	r.POST("/api/v1/task/:task_id/assign/:user_id", func(c *gin.Context) {
		taskID := c.Param("task_id")
		userID := c.Param("user_id")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := gw.tasksClient.AssignTo(ctx, &tasks.AssignToRequest{
			TaskId: taskID,
			UserId: userID,
		})

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if resp.Status.GetCode() != int32(codes.OK) {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
		return
	})

	r.GET("/api/v1/tasks", func(c *gin.Context) {

	})

	r.GET("/api/v1/task/:id", func(c *gin.Context) {

	})

	r.PUT("/api/v1/task/:id", func(c *gin.Context) {

	})

	r.DELETE("/api/v1/task/:id", func(c *gin.Context) {

	})

	r.Run()
}
