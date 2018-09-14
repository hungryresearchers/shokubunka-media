package infrastructure

import (
	"api/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	sqlHandler := NewSqlHandler()
	Migrate(sqlHandler.Conn)
	// controller
	userController := controllers.NewUserController(sqlHandler)

	// Grouping route
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// Define routes
	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })
	Router = router
}
