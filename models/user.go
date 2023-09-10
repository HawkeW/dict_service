package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(0);autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	NickName  string         `json:"nick_name"`
	Sex       int            `json:"sex"`
	LoginDate string         `json:"login_date"`
	Password  string         `json:"password"`
}

func (user *User) UpdateByStringMap(data map[string]string) {
	if data["nick_name"] != "" {
		user.NickName = data["nick_name"]
	}
	if data["phone"] != "" {
		user.Phone = data["phone"]
	}
	if data["email"] != "" {
		user.Email = data["email"]
	}
	if data["name"] != "" {
		user.Name = data["name"]
	}

	if sex, err := strconv.Atoi(data["sex"]); err == nil {
		user.Sex = sex
	}
}
