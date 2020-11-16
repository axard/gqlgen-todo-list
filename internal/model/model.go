package model

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID      int    `json:"-" db:"user_id"`
}
