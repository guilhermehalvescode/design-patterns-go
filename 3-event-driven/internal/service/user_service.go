package service

import (
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
	EventChan      chan domain.Event
}

func NewUserService(userRepository repository.UserRepository, eventChan chan domain.Event) UserService {
	return UserService{UserRepository: userRepository, EventChan: eventChan}
}

func (s *UserService) List() []domain.User {
	return s.UserRepository.List()
}

func (s *UserService) GetByID(id string) domain.User {
	return s.UserRepository.GetByID(id)
}

func (s *UserService) Create(user domain.User) domain.User {
	user = s.UserRepository.Create(user)
	s.EventChan <- domain.UserCreatedEvent{UserID: user.ID}
	return user
}

func (s *UserService) Update(id string, user domain.User) domain.User {
	newUser := s.UserRepository.Update(id, user)
	if newUser.ID != "" {
		s.EventChan <- domain.UserUpdatedEvent{UserID: newUser.ID}
	}
	return newUser
}

func (s *UserService) Delete(id string) bool {
	return s.UserRepository.Delete(id)
}
