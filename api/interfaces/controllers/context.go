package controllers

import (
	"api/domain"

	"github.com/gin-gonic/gin/binding"
)

type Context interface {
	Param(string) string
	ShouldBindWith(interface{}, binding.Binding) error
	Status(int)
	MustGet(string) interface{}
	JSON(int, interface{})
}

func CurrentUser(c Context) domain.User {
	userInterface := c.MustGet("current_user")
	currentUser := userInterface.(domain.User)
	return currentUser
}
