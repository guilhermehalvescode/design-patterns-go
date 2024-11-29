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

func (s *UserService) List() ([]domain.User, error) {
	return s.UserRepository.List(), nil
}

func (s *UserService) GetByID(id string) (domain.User, error) {
	return s.UserRepository.GetByID(id)
}

func (s *UserService) Create(user domain.User) (domain.User, error) {
	user, err := s.UserRepository.Create(user)
	if err != nil {
		return domain.User{}, err
	}
	s.EventChan <- domain.UserCreatedEvent{UserID: user.ID}
	return user, nil
}

func (s *UserService) Update(id string, user domain.User) (domain.User, error) {
	newUser, err := s.UserRepository.Update(id, user)
	if err != nil {
		return domain.User{}, err
	}
	s.EventChan <- domain.UserUpdatedEvent{UserID: newUser.ID}
	return newUser, nil
}

func (s *UserService) Delete(id string) (bool, error) {
	deleted, err := s.UserRepository.Delete(id)
	if err != nil {
		return false, err
	}
	return deleted, nil
}
