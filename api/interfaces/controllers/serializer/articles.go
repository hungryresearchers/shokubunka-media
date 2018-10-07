package serializer

import "api/domain"

type AllArticle struct {
	Articles []domain.Article `json:"articles"`
}
