package usecase

import "api/domain"

type ArticleUsecase struct {
	ArticleRepository ArticleRepository
}

func (usecase *ArticleUsecase) Create(article *domain.Article) error {
	err := usecase.ArticleRepository.Create(article)
	return err
}

func (usecase *ArticleUsecase) FetchAll(articles *[]domain.Article) error {
	err := usecase.ArticleRepository.FindAll(articles)
	return err
}
