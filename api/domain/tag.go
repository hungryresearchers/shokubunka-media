package domain

import "time"

type Tag struct {
	ID        int        `json:"tag_id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string     `json:"tag_name"`
	Articles  []*Article `gorm:"many2many:article_tags;"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp"`
}
