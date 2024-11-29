package service

import (
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/repository"
)

type NotificationService struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationService(notificationRepository repository.NotificationRepository) NotificationService {
	return NotificationService{notificationRepository: notificationRepository}
}

func (s *NotificationService) List() ([]domain.Notification, error) {
	return s.notificationRepository.List()
}

func (s *NotificationService) GetByID(id string) (domain.Notification, error) {
	return s.notificationRepository.GetByID(id)
}

func (s *NotificationService) Create(notification domain.Notification) (domain.Notification, error) {
	return s.notificationRepository.Create(notification)
}

func (s *NotificationService) Update(id string, notification domain.Notification) (domain.Notification, error) {
	return s.notificationRepository.Update(id, notification)
}

func (s *NotificationService) Delete(id string) (bool, error) {
	return s.notificationRepository.Delete(id)
}
