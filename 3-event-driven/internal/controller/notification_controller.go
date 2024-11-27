package controller

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/schemas"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/service"
)

type NotificationController struct {
	NotificationService service.NotificationService
}

func (nc *NotificationController) ListenEvents(eventsChan <-chan domain.Event) {
	slog.Info("Listening for events")
	for event := range eventsChan {
		slog.Info(fmt.Sprintf("Received event: %+v", event))
		nc.NotificationService.Create(domain.EventToNotification(event))
	}
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

func (nc *NotificationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if nc.NotificationService.Delete(id) {
		c.JSON(200, schemas.APIResponse[any]{Message: "Notification deleted", Timestamp: time.Now().Unix()})
		return
	}

	c.JSON(404, schemas.APIResponse[any]{Message: "Notification not found", Timestamp: time.Now().Unix()})
}
