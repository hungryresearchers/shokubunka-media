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
	u := domain.User{}
	c.Bind(&u)
	if err := controller.Usecase.Create(u); err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.Status(201)
}
