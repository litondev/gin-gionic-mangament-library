package models

import "time"

type Book struct {
	ID			uint 	`gorm:"primaryKey;autoIncrement"`
	Name 		string 	`gorm:"type=string;size:50;not null"`	
	Status 		string 	`gorm:"type:enum('on borrow','off borrow');default:on borrow"`
	CreatedAt  	time.Time
	UpdatedAt 	time.Time
}