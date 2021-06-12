package models

import "time"

type GuestBook struct {
	ID			uint 	`gorm:"primaryKey;autoIncrement"`
	UserId		uint
	Description	string
	CreatedAt  	time.Time
	UpdatedAt 	time.Time
	User  User `gorm:"references:ID"`
}