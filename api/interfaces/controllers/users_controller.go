package controllers

import (
	"api/domain"
	"api/interfaces/database"
	"api/service"
	"api/usecase"

	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Usecase: usecase.UserUsecase{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context, b binding.Binding) {
	u := &domain.User{}
	if err := c.ShouldBindWith(u, b); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	if err := controller.Usecase.Create(u); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	c.JSON(201, u)
}

func (controller *UserController) SignIn(c Context, b binding.Binding) {
	u := &domain.User{}
	if err := c.ShouldBindWith(u, b); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	u.EncryptPassword()
	if err := controller.Usecase.UserRepository.Find(u); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	service.SessionSet(c, u.ID)
	c.JSON(200, u)
}
