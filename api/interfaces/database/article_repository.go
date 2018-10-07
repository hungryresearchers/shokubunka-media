package database

import (
	"api/domain"

	"github.com/jinzhu/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func (repo *ArticleRepository) Create(article *domain.Article) error {
	err := repo.DB.Create(article).Error
	return err
}
