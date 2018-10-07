package usecase

import (
	"api/domain"
)

type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) Create(u *domain.User) error {
	u.Initialize()
	err := usecase.UserRepository.Create(u)
	return err
}

func (usecase *UserUsecase) Find(u *domain.User) error {
	err := usecase.UserRepository.Find(u)
	return err
}
