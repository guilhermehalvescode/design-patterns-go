package repository

import "github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/domain"

type NotificationRepository struct {
	Notifications []domain.Notification
}

func NewNotificationRepository() NotificationRepository {
	return NotificationRepository{Notifications: []domain.Notification{}}
}

func (r *NotificationRepository) List() []domain.Notification {
	return r.Notifications
}

func (r *NotificationRepository) GetByID(id string) domain.Notification {
	for _, notification := range r.Notifications {
		if notification.ID == id {
			return notification
		}
	}
	return domain.Notification{}
}

func (r *NotificationRepository) Create(notification domain.Notification) domain.Notification {
	r.Notifications = append(r.Notifications, notification)
	return notification
}

func (r *NotificationRepository) Update(id string, notification domain.Notification) domain.Notification {
	for i, n := range r.Notifications {
		if n.ID == id {
			r.Notifications[i] = notification
			return notification
		}
	}
	return domain.Notification{}
}

func (r *NotificationRepository) Delete(id string) bool {
	for i, notification := range r.Notifications {
		if notification.ID == id {
			r.Notifications = append(r.Notifications[:i], r.Notifications[i+1:]...)
			return true
		}
	}
	return false
}
