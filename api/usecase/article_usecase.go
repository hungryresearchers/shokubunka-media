package usecase

import "api/domain"

type ArticleUsecase struct {
	ArticleRepository ArticleRepository
}

func (usecase *ArticleUsecase) Create(article *domain.Article) error {
	err := usecase.ArticleRepository.Create(article)
	return err
}
