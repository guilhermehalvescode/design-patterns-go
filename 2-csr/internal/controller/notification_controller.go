package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/schemas"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/service"
)

type NotificationController struct {
	NotificationService service.NotificationService
}

func (nc *NotificationController) List(c *gin.Context) {
	notifications := nc.NotificationService.List()
	c.JSON(200, schemas.APIResponse[[]domain.Notification]{Message: "List of notifications", Timestamp: time.Now().Unix(), Data: notifications})
}

func (nc *NotificationController) GetByID(c *gin.Context) {
	notification := nc.NotificationService.GetByID(c.Param("id"))
	if notification.ID != "" {
		c.JSON(200, schemas.APIResponse[domain.Notification]{Message: "Notification found", Timestamp: time.Now().Unix(), Data: notification})
		return
	}
	c.JSON(404, schemas.APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}

func (nc *NotificationController) Create(c *gin.Context) {
	var notification domain.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(400, schemas.APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	notification = nc.NotificationService.Create(notification)
	c.JSON(201, schemas.APIResponse[domain.Notification]{Message: "Notification created", Timestamp: time.Now().Unix(), Data: notification})
}

func (nc *NotificationController) Update(c *gin.Context) {
	id := c.Param("id")
	var notification domain.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(400, schemas.APIResponse[any]{Message: "Invalid request", Timestamp: time.Now().Unix()})
		return
	}
	notification = nc.NotificationService.Update(id, notification)
	if notification.ID != "" {
		c.JSON(200, schemas.APIResponse[domain.Notification]{Message: "Notification updated successfully", Timestamp: time.Now().Unix(), Data: notification})
		return
	}
	c.JSON(404, schemas.APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}

func (nc *NotificationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if nc.NotificationService.Delete(id) {
		c.JSON(200, schemas.APIResponse[any]{Message: "Notification deleted", Timestamp: time.Now().Unix()})
		return
	}

	c.JSON(404, schemas.APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}
