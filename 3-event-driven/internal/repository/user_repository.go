package repository

import (
	"errors"

	"github.com/guilhermehalvescode/design-patterns-go/3-event-driven/internal/domain"
)

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() UserRepository {
	return UserRepository{users: []domain.User{}}
}

func (r *UserRepository) List() []domain.User {
	return r.users
}

func (r *UserRepository) GetByID(id string) (domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepository) Update(id string, user domain.User) (domain.User, error) {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i] = user
			return user, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (r *UserRepository) Delete(id string) (bool, error) {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("user not found")
}
