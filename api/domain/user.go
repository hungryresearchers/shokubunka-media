package domain

import (
	"api/service"
	"time"
)

type User struct {
	ID                 int    `json:"user_id" gorm:"primary_key;AUTO_INCREMENT" sql:"not null" json:"-"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	NickName           string `json:"nickname"`
	Email              string `sql:"unique_index;not null" json:"email,omitempty" binding:"exists,email" valid:"email"`
	EncryptedPassword  string `sql:"not null" json:"password,omitempty" valid:length(8:255)`
	InvitationToken    string `sql:"unique_index;not null" json:"invitation_token,omitempty"`
	ResetPasswordToken string `json:"reset_password_token,omitempty"`
	Role               *int   `sql:"default:0" json:"role,omitempty"` // role 0: user, 1: author, 2: admin
	Articles           []Article
	CreatedAt          time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"type:timestamp"`
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
