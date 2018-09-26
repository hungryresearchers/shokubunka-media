package controllers

import (
	"api/domain"
	"api/interfaces/database"
	"api/usecase"

	"github.com/gin-gonic/gin/binding"
)

type ArticleController struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
	return &ArticleController{
		Usecase: usecase.ArticleUsecase{
			ArticleRepository: &database.ArticleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArticleController) Create(c Context, b binding.Binding) {
	article := &domain.Article{}
	if err := c.ShouldBindWith(article, b); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	if err := controller.Usecase.Create(article); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	c.JSON(201, article)
}
