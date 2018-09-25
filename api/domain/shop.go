package domain

import "time"

type Shop struct {
	ID            int        `json:"shop_id" gorm:"primary_key;AUTO_INCREMENT"`
	Name          string     `json:"shop_name"`
	MinPrice      int        `json:"min_price"`
	MaxPrice      int        `json:"max_price"`
	LocationURL   string     `json:"location_url"`
	ContactNumber string     `json:"contact_number"`
	OpenAt        time.Time  `json:"open_at"`
	ClosedAt      time.Time  `json:"closed_at"`
	SnsURL        string     `json:"sns_url"`
	Articles      []*Article `gorm:"many2many:"article_shops;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
