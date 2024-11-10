package domain

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	UserID int    `json:"user_id"`
}
