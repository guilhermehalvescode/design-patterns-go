package domain

import "github.com/google/uuid"

type Event interface {
	GetMessage() string
	GetUserID() string
}

type Notification struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

func EventToNotification(event Event) Notification {
	return Notification{
		ID:      uuid.New().String(),
		Message: event.GetMessage(),
		UserID:  event.GetUserID(),
	}
}
