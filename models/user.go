package models

import "time"

type User struct {
	ID			uint 	`gorm:"primaryKey;autoIncrement"`
	Username 	string 	`gorm:"type=string;size:50;not null;index"`
	Email 		string 	`gorm:"type=string;size:50;not null;unique"`
	Password 	string 	`gorm:"not null"`
	Role 		string 	`gorm:"type:enum('user','admin');default:user"`
	Photo 		string 	`gorm:"default:default.png;size:50"`
	CreatedAt  	time.Time
	UpdatedAt 	time.Time
}