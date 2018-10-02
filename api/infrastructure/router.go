package infrastructure

import (
	"api/domain"
	"api/interfaces/controllers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qor/admin"
	"github.com/qor/validations"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	mux := http.NewServeMux()

	sqlHandler := NewSqlHandler()
	Migrate(sqlHandler.Conn)
	validations.RegisterCallbacks(sqlHandler.Conn)

	// Admin config & routing
	Admin := admin.New(&admin.AdminConfig{
		DB:       sqlHandler.Conn,
		SiteName: "Hungry Researchers",
	})
	user := Admin.AddResource(&domain.User{})
	AddResourceValidator(user)
	defineUserMetaInfo(user)
	Admin.MountTo("/admin", mux)
	router.Any("/admin/*resources", gin.WrapH(mux))

	// Session Setting
	store, _ := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("shokubunka_session", store))

	// Grouping route
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// Users
	users := v1.Group("/users")
	userController := controllers.NewUserController(sqlHandler)
	users.POST("", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		userController.Create(c, b)
	})
	users.POST("/login", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		userController.SignIn(c, b)
	})

	// Articles
	articles := v1.Group("/articles")
	articleController := controllers.NewArticleController(sqlHandler)
	articles.POST("", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		articleController.Create(c, b)
	})
	Router = router
}
