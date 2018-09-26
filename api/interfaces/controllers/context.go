package controllers

import "github.com/gin-gonic/gin/binding"

type Context interface {
	Param(string) string
	ShouldBindWith(interface{}, binding.Binding) error
	Status(int)
	JSON(int, interface{})
}
