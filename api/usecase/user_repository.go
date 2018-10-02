package usecase

import "api/domain"

type UserRepository interface {
	Create(*domain.User) error
	Find(*domain.User) error
}
