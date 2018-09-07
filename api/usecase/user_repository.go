package usecase

import "api/domain"

type UserRepository interface {
	Create(domain.User) domain.User
}
