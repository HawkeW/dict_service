package models

type WordBook struct {
	Id          uint   `json:"id"`
	UserId      int    `json:"user_id"`
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
