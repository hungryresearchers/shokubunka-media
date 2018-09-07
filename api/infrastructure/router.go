package infrastructure

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	router := gin.Default()
	// controller
	userController := controllers.NewUserController(NewSqlHandler)

	// Define routes
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	Router = router
}
