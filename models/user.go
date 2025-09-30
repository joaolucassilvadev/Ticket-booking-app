package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  = UserRole("manager")
	Attendee = UserRole("attendee")
)

type User struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	Email    string    `json:"email" gorm:"unique"`
	Role     UserRole  `json:"role" gorm:"text;default:'attendee'"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
}
