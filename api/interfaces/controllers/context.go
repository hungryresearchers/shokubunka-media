package controllers

import (
	"api/domain"
	"mime/multipart"

	"github.com/gin-gonic/gin/binding"
)

type Context interface {
	Param(string) string
	ShouldBindWith(interface{}, binding.Binding) error
	FormFile(string) (*multipart.FileHeader, error)
	AbortWithStatus(int)
	Status(int)
	MustGet(string) interface{}
	JSON(int, interface{})
}

func CurrentUser(c Context) domain.User {
	userInterface := c.MustGet("current_user")
	currentUser := userInterface.(domain.User)
	return currentUser
}
