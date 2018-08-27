package database

import "shokubunka-media/api/src/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Create(u *domain.User) (Result, error) {
	result, err := repo.Create(u)
	return result, err
}
