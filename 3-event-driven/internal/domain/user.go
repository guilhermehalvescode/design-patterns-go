package domain

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserCreatedEvent struct {
	UserID string
}

func (e UserCreatedEvent) GetMessage() string {
	return "User created"
}

func (e UserCreatedEvent) GetUserID() string {
	return e.UserID
}

type UserUpdatedEvent struct {
	UserID string
}

func (e UserUpdatedEvent) GetMessage() string {
	return "User updated"
}

func (e UserUpdatedEvent) GetUserID() string {
	return e.UserID
}
