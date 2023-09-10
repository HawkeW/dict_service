package models

import (
	"time"
)

type WordBook struct {
	Id          uint
	UserId      int
	Type        int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     time.Time
}
