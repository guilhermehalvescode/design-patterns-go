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
	notifications, err := nc.NotificationService.List()
	if err != nil {
		c.JSON(500, schemas.NewAPIMessageResponse("Internal server error"))
		return
	}
	c.JSON(200, schemas.NewAPIResponse("List of notifications", notifications))
}

func (nc *NotificationController) GetByID(c *gin.Context) {
	notification, err := nc.NotificationService.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(404, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(200, schemas.NewAPIResponse("Notification found", notification))
}

func (nc *NotificationController) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := nc.NotificationService.Delete(id)
	if err != nil {
		c.JSON(404, schemas.NewAPIMessageResponse(err.Error()))
		return
	}
	c.JSON(200, schemas.NewAPIMessageResponse("Notification deleted"))
}
