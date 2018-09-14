package database

import "api/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Create(u *domain.User) error {
	err := repo.SqlHandler.Create(u)
	return err
}
