package domain

import (
	"api/service"
	"time"
)

type User struct {
	ID                 int    `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	FirstName          string `json:"firstname" gorm:"not null" valid:"required"`
	LastName           string `json:"lastname" gorm:"not null" valid:"required"`
	NickName           string `json:"nickname"`
	Email              string `json:"email,omitempty" gorm:"unique_index;not null" binding:"exists,email" valid:"email,required"`
	EncryptedPassword  string `json:"password,omitempty" gorm:"not null" valid:"required,length(8:255)"`
	InvitationToken    string `json:"invitation_token,omitempty gorm:"unique_index;not null" valid:"required"`
	ResetPasswordToken string `json:"reset_password_token,omitempty"`
	Role               *int   `json:"role,omitempty" gorm:"default:0;not null" valid:"required"` // role 0: user, 1: author, 2: admin
	Articles           []Article
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
