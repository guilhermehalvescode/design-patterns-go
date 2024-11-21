package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/service"
)

type NotificationController struct {
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) NotificationController {
	return NotificationController{notificationService: notificationService}
}

func (nc *NotificationController) List(c *gin.Context) {
	notifications := nc.notificationService.List()
	c.JSON(200, APIResponse[[]domain.Notification]{Message: "List of notifications", Timestamp: time.Now().Unix(), Data: notifications})
}

func (nc *NotificationController) GetByID(c *gin.Context) {
	notification := nc.notificationService.GetByID(c.Param("id"))
	if notification.ID != "" {
		c.JSON(200, APIResponse[domain.Notification]{Message: "Notification found", Timestamp: time.Now().Unix(), Data: notification})
		return
	}
	c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}

func (nc *NotificationController) Create(c *gin.Context) {
	var notification domain.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	notification = nc.notificationService.Create(notification)
	c.JSON(201, APIResponse[domain.Notification]{Message: "Notification created", Timestamp: time.Now().Unix(), Data: notification})
}

func (nc *NotificationController) Update(c *gin.Context) {
	id := c.Param("id")
	var notification domain.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(400, APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	notification = nc.notificationService.Update(id, notification)
	if notification.ID != "" {
		c.JSON(200, APIResponse[domain.Notification]{Message: "Notification updated successfully", Timestamp: time.Now().Unix(), Data: notification})
		return
	}
	c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}

func (nc *NotificationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if nc.notificationService.Delete(id) {
		c.JSON(200, APIResponse[any]{Message: "Notification deleted", Timestamp: time.Now().Unix()})
		return
	}

	c.JSON(404, APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}
