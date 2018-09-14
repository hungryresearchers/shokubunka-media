package controllers

import (
	"api/domain"
	"api/interfaces/database"
	"api/usecase"
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

func (controller *UserController) Create(c Context) {
	u := &domain.User{}
	if err := c.Bind(u); err != nil {
		c.JSON(400, NewError(err))
	}
	if err := controller.Usecase.Create(u); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	c.Status(201)
}
