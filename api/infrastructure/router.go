package infrastructure

import (
	"api/domain"
	"api/interfaces/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	mux := http.NewServeMux()

	sqlHandler := NewSqlHandler()
	Migrate(sqlHandler.Conn)

	// Admin config & routing
	Admin := admin.New(&admin.AdminConfig{DB: sqlHandler.Conn})
	Admin.AddResource(&domain.User{})
	Admin.MountTo("/admin", mux)
	router.Any("/admin/*resources", gin.WrapH(mux))

	// controller
	userController := controllers.NewUserController(sqlHandler)

	// Grouping route
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// Define routes
	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })
	Router = router
}
