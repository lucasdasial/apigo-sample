package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
