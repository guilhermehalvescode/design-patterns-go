package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse[K any] struct {
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      K      `json:"data,omitempty"`
}

type Notification struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	engine := gin.Default()

	users := []User{}

	notifications := []Notification{}

	// User routes
	engine.GET("/users", func(c *gin.Context) {
		c.JSON(200, APIResponse[[]User]{Message: "List of users", Timestamp: time.Now().Unix(), Data: users})
	})

	engine.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if user.ID == id {
				c.JSON(200, APIResponse[User]{Message: "User found", Timestamp: time.Now().Unix(), Data: user})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "User not found", Timestamp: time.Now().Unix()})
	})

	engine.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		users = append(users, user)
		c.JSON(201, APIResponse[User]{Message: "User created", Timestamp: time.Now().Unix(), Data: user})
	})

	engine.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		for i, u := range users {
			if u.ID == id {
				users[i] = user
				c.JSON(200, APIResponse[User]{Message: "User updated successfully", Timestamp: time.Now().Unix(), Data: user})
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

	// Notification routes
	engine.GET("/notifications", func(c *gin.Context) {
		c.JSON(200, APIResponse[[]Notification]{Message: "List of notifications", Timestamp: time.Now().Unix(), Data: notifications})
	})

	engine.GET("/notifications/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, notification := range notifications {
			if notification.ID == id {
				c.JSON(200, APIResponse[Notification]{Message: "Notification found", Timestamp: time.Now().Unix(), Data: notification})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
	})

	engine.POST("/notifications", func(c *gin.Context) {
		var notification Notification
		if err := c.ShouldBindJSON(&notification); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		notification.ID = string(len(notifications) + 1)
		notifications = append(notifications, notification)
		c.JSON(201, APIResponse[Notification]{Message: "Notification created", Timestamp: time.Now().Unix(), Data: notification})
	})

	engine.PUT("/notifications/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedNotification Notification
		if err := c.ShouldBindJSON(&updatedNotification); err != nil {
			c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
			return
		}
		for i, notification := range notifications {
			if notification.ID == id {
				notifications[i] = updatedNotification
				notifications[i].ID = id
				c.JSON(200, APIResponse[Notification]{Message: "Notification updated successfully", Timestamp: time.Now().Unix(), Data: notifications[i]})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
	})

	engine.DELETE("/notifications/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, notification := range notifications {
			if notification.ID == id {
				notifications = append(notifications[:i], notifications[i+1:]...)
				c.JSON(200, APIResponse[any]{Message: "Notification deleted successfully", Timestamp: time.Now().Unix()})
				return
			}
		}
		c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
	})

	engine.Run(":8080")
}
