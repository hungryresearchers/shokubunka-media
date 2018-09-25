package infrastructure

import (
	"api/domain"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Article{})
	db.AutoMigrate(&domain.Tag{})
	db.AutoMigrate(&domain.Shop{})
}
