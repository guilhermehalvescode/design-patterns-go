package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/schemas"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/service"
)

type NotificationController struct {
	NotificationService service.NotificationService
}

func (nc *NotificationController) ListenEvents(eventsChan <-chan domain.Event) {
	for event := range eventsChan {
		nc.NotificationService.Create(domain.EventToNotification(event))
	}
}

func (nc *NotificationController) List(c *gin.Context) {
	notifications := nc.NotificationService.List()
	c.JSON(200, schemas.NewAPIResponse("List of notifications", notifications))
}

func (nc *NotificationController) GetByID(c *gin.Context) {
	notification := nc.NotificationService.GetByID(c.Param("id"))
	if notification.ID != "" {
		c.JSON(200, schemas.NewAPIResponse("Notification found", notification))
		return
	}
	c.JSON(404, schemas.NewAPIMessageResponse("Notification not found"))
}

func (nc *NotificationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if nc.NotificationService.Delete(id) {
		c.JSON(200, schemas.NewAPIMessageResponse("Notification deleted"))
		return
	}

	c.JSON(404, schemas.NewAPIMessageResponse("Notification not found"))
}
