package database

import "api/domain"

type ArticleRepository struct {
	SqlHandler
}

func (repo *ArticleRepository) Create(article *domain.Article) error {
	err := repo.SqlHandler.Create(article)
	return err
}
