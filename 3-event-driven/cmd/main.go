package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/controller"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/repository"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/service"
)

func main() {
	engine := gin.Default()

	eventsChan := make(chan domain.Event)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, eventsChan)
	userController := controller.UserController{
		UserService: userService,
	}

	// User routes
	engine.GET("/users", userController.List)
	engine.GET("/users/:id", userController.GetByID)
	engine.POST("/users", userController.Create)
	engine.PUT("/users/:id", userController.Update)
	engine.DELETE("/users/:id", userController.Delete)

	notificationRepository := repository.NewNotificationRepository()
	notificationService := service.NewNotificationService(notificationRepository)

	notificationController := controller.NotificationController{
		NotificationService: notificationService,
	}
	// Notification routes
	engine.GET("/notifications", notificationController.List)
	engine.GET("/notifications/:id", notificationController.GetByID)

	engine.DELETE("/notifications/:id", notificationController.Delete)

	go notificationController.ListenEvents(eventsChan)

	engine.Run(":8080")
}
