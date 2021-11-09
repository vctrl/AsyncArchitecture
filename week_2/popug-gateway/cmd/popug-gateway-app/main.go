package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type gateway struct {
	// auth
	// tasks
	// logger
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// AuthService.CheckSession()
		token, err := c.Cookie("session_token")
		if err != nil {

		}

		fmt.Println(token) // todo check session token
	}
}

func main() {
	r := gin.Default()

	// create auth service client
	// create tasks service client
	r.Use(SessionMiddleware())

	// auth
	r.POST("/api/v1/login", func(c *gin.Context) {
		// AuthServiceClient.Login()
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
