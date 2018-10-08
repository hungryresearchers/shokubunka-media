package usecase

import "api/domain"

type ArticleRepository interface {
	Create(*domain.Article) error
	FindAll(*[]domain.Article) error
	Find(*domain.Article) error
	Destroy(*domain.Article) error
}
