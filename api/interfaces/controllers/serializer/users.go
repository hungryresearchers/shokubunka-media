package serializer

import "time"

type UserResponse struct {
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	NickName  string    `json:"nickname",omitempty`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
