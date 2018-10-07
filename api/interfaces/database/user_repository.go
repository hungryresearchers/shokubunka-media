package database

import (
	"api/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Create(u *domain.User) error {
	err := repo.DB.Create(u).Error
	return err
}

func (repo *UserRepository) Find(u *domain.User) error {
	err := repo.DB.Where(u).First(u).Error
	return err
}
