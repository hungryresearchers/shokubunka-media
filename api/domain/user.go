package domain

import (
	"api/service"
)

type User struct {
	ID                 int    `gorm:"primary_key;AUTO_INCREMENT" sql:"not null" json:"-"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	NickName           string `json:"nickname"`
	Email              string `sql:"unique_index;not null" json:"email,omitempty"`
	EncryptedPassword  string `sql:"not null" json:"password,omitempty"`
	InvitationToken    string `sql:"unique_index;not null" json:"invitation_token,omitempty"`
	ResetPasswordToken string `json:"reset_password_token,omitempty"`
	Role               *int   `sql:"default:0" json:"role,omitempty"` // role 0: user, 1: author, 2: admin
}

func (u *User) Initialize() {
	u.InvitationToken = service.GenerateToken()
	u.EncryptPassword()
}

func (u *User) EncryptPassword() {
	unsafePassword := u.EncryptedPassword
	u.EncryptedPassword = service.ToHash(unsafePassword)
}
