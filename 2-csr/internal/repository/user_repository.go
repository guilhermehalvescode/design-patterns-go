package repository

import "github.com/guilhermehalvescode/design-patterns-go/2-csr/internal/domain"

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() UserRepository {
	return UserRepository{users: []domain.User{}}
}

func (r *UserRepository) List() []domain.User {
	return r.users
}

func (r *UserRepository) GetByID(id string) domain.User {
	for _, user := range r.users {
		if user.ID == id {
			return user
		}
	}
	return domain.User{}
}

func (r *UserRepository) Create(user domain.User) domain.User {
	r.users = append(r.users, user)
	return user
}

func (r *UserRepository) Update(id string, user domain.User) domain.User {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i] = user
			return user
		}
	}
	return domain.User{}
}

func (r *UserRepository) Delete(id string) bool {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true
		}
	}
	return false
}
