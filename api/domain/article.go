package domain

import "time"

type Article struct {
	ID           int `json:"article_id" gorm:"primary_key,AUTO_INCREMENT"`
	Title        string
	Body         string `sql:"type:text"`
	ThumbnailURL string `json:"thumbnail_url"`
	UserID       int
	Tags         []*Tag  `gorm:"many2many:article_tags;"`
	Shops        []*Shop `gorm:"many2many:article_shops;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
