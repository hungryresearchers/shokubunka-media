package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName          string
	LastName           string
	NickName           string
	Email              string `gorm:"unique_index;not null"`
	EncryptedPassword  string `gorm:"not null"`
	InvitationToken    string `gorm:"unique_index;not null"`
	ResetPasswordToken string
	Role               int `gorm:"not null"` // role 0: user, 1: author, 2: admin
}
