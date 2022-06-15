package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

func NewUser(name string, age int16) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func NewUserId(id uint, name string, age int16) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}
