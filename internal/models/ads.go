package models

type Ad struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserId      int    `json:"user_id"`
}
