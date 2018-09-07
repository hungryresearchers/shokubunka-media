package usecase

import "api/domain"

type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) Create(u *domain.User) (user *interface{}, err error) {
	user, err = usecase.UserRepository.Create(u)
	return
}
