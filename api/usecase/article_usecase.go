package usecase

import (
	"api/domain"
	"errors"
)

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

func (usecase *ArticleUsecase) Destroy(article *domain.Article, user *domain.User) error {
	if err := usecase.ArticleRepository.Find(article); err != nil {
		return err
	}
	if user.ID != article.UserID && user.Role != 2 {
		return errors.New("Permission error: Request not permitted")
	}
	if err := usecase.ArticleRepository.Destroy(article); err != nil {
		return err
	}
	return nil
}
