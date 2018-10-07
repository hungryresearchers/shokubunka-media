package controllers

import (
	"api/domain"
	"api/interfaces/controllers/serializer"
	"api/interfaces/database"
	"api/usecase"

	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type ArticleController struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleController(sqlhandler *gorm.DB) *ArticleController {
	return &ArticleController{
		Usecase: usecase.ArticleUsecase{
			ArticleRepository: &database.ArticleRepository{
				DB: sqlhandler,
			},
		},
	}
}

func (controller *ArticleController) Create(c Context, b binding.Binding) {
	user := CurrentUser(c)
	article := &domain.Article{UserID: user.ID}
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

func (controller *ArticleController) Index(c Context) {
	articles := make([]domain.Article, 100)
	if err := controller.Usecase.FetchAll(&articles); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	c.JSON(200, serializer.AllArticle{Articles: articles})
}
