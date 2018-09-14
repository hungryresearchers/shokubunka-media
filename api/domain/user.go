package domain

type User struct {
	ID                 int    `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	FirstName          string `gorm:"not null" json:"firstname"`
	LastName           string `gorm:"not null" json:"lastname"`
	NickName           string `gorm:"not null" json:"nickname"`
	Email              string `gorm:"unique_index;not null" json:"email,omitempty"`
	EncryptedPassword  string `gorm:"not null" json:"-"`
	InvitationToken    string `gorm:"unique_index;not null" json:"invitation_token,omitempty"`
	ResetPasswordToken string `gorm:"unique_index" json:"reset_password_token,omitempty"`
	Role               int    `gorm:"not null" json:"role,omitempty"` // role 0: user, 1: author, 2: admin
}
