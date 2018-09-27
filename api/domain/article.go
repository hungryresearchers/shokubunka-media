package domain

import "time"

type Article struct {
	ID           int       `json:"-" gorm:"primary_key,AUTO_INCREMENT"`
	Title        string    `json:"title" sql:"not null"`
	Body         string    `json:"body" sql:"not null;type:text"`
	ThumbnailURL string    `json:"thumbnail_url" sql:"not null"`
	UserID       int       `json:"-" sql:"not null"`
	Tags         []*Tag    `json:"tags,omitempty" gorm:"many2many:article_tags;"`
	Shops        []*Shop   `json:"shops,omitempty" gorm:"many2many:article_shops;"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"type:timestamp" sql:"DEFAULT:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"type:timestamp" sql:"DEFAULT:current_timestamp"`
}
