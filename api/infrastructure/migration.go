package infrastructure

import (
	"api/domain"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}
