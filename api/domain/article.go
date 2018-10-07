package domain

import "time"

type Article struct {
	ID           int       `json:"id" gorm:"primary_key,AUTO_INCREMENT;not null"`
	Title        string    `json:"title" gorm:"not null"`
	Body         string    `json:"body" gorm:"not null;type:text"`
	ThumbnailURL string    `json:"thumbnail_url" gorm:"column:thumbnail_url;not null"`
	UserID       int       `json:"user_id" gorm:"not null;column:user_id"`
	Tags         []*Tag    `json:"tags,omitempty" gorm:"many2many:article_tags;"`
	Shops        []*Shop   `json:"shops,omitempty" gorm:"many2many:article_shops;"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"type:timestamp;DEFAULT:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"type:timestamp;DEFAULT:current_timestamp"`
}
