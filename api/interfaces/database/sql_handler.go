package database

import "api/domain"

type SqlHandler interface {
	Create(*domain.User) error
}
