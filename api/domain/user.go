package domain

import (
	"api/service"
)

type User struct {
	ID                 int    `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	NickName           string `json:"nickname"`
	Email              string `gorm:"unique_index;not null" json:"email,omitempty"`
	EncryptedPassword  string `gorm:"not null" json:"password"`
	InvitationToken    string `gorm:"unique_index;not null" json:"invitation_token,omitempty"`
	ResetPasswordToken string `gorm:"unique_index" json:"reset_password_token,omitempty"`
	Role               int    `gorm:"not null" json:"role,omitempty"` // role 0: user, 1: author, 2: admin
}

func (u *User) EncryptPassword() {
	unsafePassword := u.EncryptedPassword
	u.EncryptedPassword = service.ToHash(unsafePassword)
}
