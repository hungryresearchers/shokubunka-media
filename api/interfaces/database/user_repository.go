package database

import "api/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Create(u *domain.User) (*interface{}, error) {
	result, err := repo.Create(u)
	return result, err
}
