package database

import (
	"api/domain"

	"github.com/jinzhu/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

// Create #Create article
func (repo *ArticleRepository) Create(article *domain.Article) error {
	err := repo.DB.Create(article).Error
	return err
}

// FindAll #Find all aritcle and order created_at to desc
func (repo *ArticleRepository) FindAll(articles *[]domain.Article) error {
	err := repo.DB.Order("created_at desc").Limit(100).Find(articles).Error
	return err
}

// Find #Find only one article
func (repo *ArticleRepository) Find(article *domain.Article) error {
	err := repo.DB.Where(article).First(article).Error
	return err
}

// Destroy #Destroy article
func (repo *ArticleRepository) Destroy(article *domain.Article) error {
	err := repo.DB.Delete(article).Error
	return err
}
