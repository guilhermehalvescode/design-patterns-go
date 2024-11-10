package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/internal/adhoc/domain"
)

type APIResponse[K any] struct {
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      K      `json:"data,omitempty"`
}

func main() {
	engine := gin.Default()

	users := []domain.User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
		{ID: "4", Name: "David"},
		{ID: "5", Name: "Eve"},
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, APIResponse[string]{Message: "pong", Timestamp: time.Now().Unix(), Data: "pong"})
	})

	// User routes
	engine.GET("/users", func(c *gin.Context) {
		c.JSON(200, APIResponse[[]domain.User]{Message: "List of users", Timestamp: time.Now().Unix(), Data: users})
	})

	engine.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if user.ID == id {
				c.JSON(200, APIResponse[domain.User]{Message: "User found", Timestamp: time.Now().Unix(), Data: user})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
	})

	engine.POST("/users", func(c *gin.Context) {
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		users = append(users, user)
		c.JSON(201, APIResponse[domain.User]{Message: "User created", Timestamp: time.Now().Unix(), Data: user})
	})

	engine.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		for i, u := range users {
			if u.ID == id {
				users[i] = user
				c.JSON(200, APIResponse[domain.User]{Message: "User updated successfully", Timestamp: time.Now().Unix(), Data: user})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
	})

	engine.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(200, APIResponse[any]{Message: "User deleted successfully", Timestamp: time.Now().Unix()})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
	})

	engine.Run(":8080")
}
