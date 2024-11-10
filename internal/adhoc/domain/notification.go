package domain

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}
