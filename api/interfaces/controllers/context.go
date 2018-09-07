package controllers

type Context {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}