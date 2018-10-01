package domain

import "time"

type Tag struct {
	ID        int        `json:"tag_id" gorm:"primary_key;AUTO_INCREMENT;not null"`
	Name      string     `json:"tag_name" gorm:"not null"`
	Articles  []*Article `gorm:"many2many:article_tags;"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;DEFAULT:current_timestamp;not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp;DEFAULT:current_timestamp;not null"`
}
