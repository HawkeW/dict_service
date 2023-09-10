package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Name      sql.NullString
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	NickName  string
	Sex       int
	LoginDate time.Time
}
