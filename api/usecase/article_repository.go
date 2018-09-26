package usecase

import "api/domain"

type ArticleRepository interface {
	Create(*domain.Article) error
}
