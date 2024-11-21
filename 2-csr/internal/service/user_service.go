package service

import (
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/repository"
)

type UserService struct {
	UserRepository      repository.UserRepository
	NotificationService NotificationService
}

func NewUserService(userRepository repository.UserRepository, notificationService NotificationService) UserService {
	return UserService{UserRepository: userRepository, NotificationService: notificationService}
}

func (s *UserService) List() []domain.User {
	return s.UserRepository.List()
}

func (s *UserService) GetByID(id string) domain.User {
	return s.UserRepository.GetByID(id)
}

func (s *UserService) Create(user domain.User) domain.User {
	user = s.UserRepository.Create(user)
	return user
}

func (s *UserService) Update(id string, user domain.User) domain.User {
	return s.UserRepository.Update(id, user)
}

func (s *UserService) Delete(id string) bool {
	return s.UserRepository.Delete(id)
}
