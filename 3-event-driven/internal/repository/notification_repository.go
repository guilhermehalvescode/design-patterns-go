package repository

import (
	"errors"

	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
)

type NotificationRepository struct {
	Notifications []domain.Notification
}

func NewNotificationRepository() NotificationRepository {
	return NotificationRepository{Notifications: []domain.Notification{}}
}

func (r *NotificationRepository) List() ([]domain.Notification, error) {
	return r.Notifications, nil
}

func (r *NotificationRepository) GetByID(id string) (domain.Notification, error) {
	for _, notification := range r.Notifications {
		if notification.ID == id {
			return notification, nil
		}
	}
	return domain.Notification{}, errors.New("notification not found")
}

func (r *NotificationRepository) Create(notification domain.Notification) (domain.Notification, error) {
	r.Notifications = append(r.Notifications, notification)
	return notification, nil
}

func (r *NotificationRepository) Update(id string, notification domain.Notification) (domain.Notification, error) {
	for i, n := range r.Notifications {
		if n.ID == id {
			r.Notifications[i] = notification
			return notification, nil
		}
	}
	return domain.Notification{}, errors.New("notification not found")
}

func (r *NotificationRepository) Delete(id string) (bool, error) {
	for i, notification := range r.Notifications {
		if notification.ID == id {
			r.Notifications = append(r.Notifications[:i], r.Notifications[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("notification not found")
}
