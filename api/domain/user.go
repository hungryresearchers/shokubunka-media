package domain

import (
	"api/service"
	"time"
)

type User struct {
	ID                 int       `json:"-" gorm:"primary_key;AUTO_INCREMENT;not null"`
	FirstName          string    `json:"firstname,omitempty" gorm:"not null" binding:"exists"`
	LastName           string    `json:"lastname,omitempty" gorm:"not null" binding:"exists"`
	NickName           string    `json:"nickname,omitempty"`
	Email              string    `json:"email,omitempty" gorm:"unique_index;not null" binding:"exists,email"`
	EncryptedPassword  string    `json:"password,omitempty" gorm:"not null"`
	InvitationToken    string    `json:"invitation_token,omitempty" gorm:"unique_index;not null"`
	ResetPasswordToken string    `json:"reset_password_token,omitempty"`
	Role               int       `json:"role,string,omitempty" gorm:"default:0;not null"` // role 0: user, 1: author, 2: admin
	Articles           []Article `json:"articles,omitempty"`
	CreatedAt          time.Time `json:"created_at" gorm:"type:timestamp;DEFAULT:current_timestamp"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"type:timestamp;DEFAULT:current_timestamp"`
}

func (u *User) BeforeSave() (err error) {
	token, err := service.GenerateToken()
	u.InvitationToken = token
	return
}

func (u *User) Initialize() {
	u.InvitationToken, _ = service.GenerateToken()
	u.EncryptPassword()
}

func (u *User) EncryptPassword() {
	unsafePassword := u.EncryptedPassword
	u.EncryptedPassword = service.ToHash(unsafePassword)
}
