package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/controller"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/repository"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/service"
)

func main() {
	engine := gin.Default()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
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
	engine.POST("/notifications", notificationController.Create)
	engine.PUT("/notifications/:id", notificationController.Update)

	engine.DELETE("/notifications/:id", notificationController.Delete)

	engine.Run(":8080")
}
