package models

import (
	"gorm.io/gorm"
	"time"
)

type WordBook struct {
	Id          uint           `json:"id"`
	UserId      int            `json:"user_id"`
	Type        int            `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (book WordBook) FromMap(userId int, data map[string]string) WordBook {
	return WordBook{
		UserId:      userId,
		Type:        1,
		Name:        data["name"],
		Description: data["description"],
	}
}

func (book WordBook) UpdateByStringMap(data map[string]string) WordBook {
	if data["name"] != "" {
		book.Name = data["name"]
	}
	if data["description"] != "" {
		book.Description = data["description"]
	}
	return book
}
