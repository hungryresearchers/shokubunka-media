package infrastructure

import (
	"api/domain"
	"api/infrastructure/config"
	"api/infrastructure/middleware"
	"api/interfaces/controllers"
	"api/interfaces/database"
	"context"
	"log"
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
	sqlhandler := database.NewSQLHandler()

	//
	Migrate(sqlhandler)
	validations.RegisterCallbacks(sqlhandler)
	ctx := context.Background()
	blob, err := config.Setup(ctx, "gcp")
	if err != nil {
		log.Fatal(err)
	}

	// Admin config & routing
	Admin := admin.New(&admin.AdminConfig{
		DB:       sqlhandler,
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

	v1.Use(middleware.AuthMiddleware())

	// ImageUploader
	images := v1.Group("/images")
	imageController := controllers.NewImageController()
	images.Use(middleware.ResourcePermissionMiddleware())
	images.POST("/upload", func(c *gin.Context) {
		imageController.Upload(c, blob, ctx)
	})

	// Users
	users := v1.Group("/users")
	userController := controllers.NewUserController(sqlhandler)
	users.POST("/login", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		userController.SignIn(c, b)
	})
	users.Use(middleware.UserPermissionMiddleware())
	users.GET("/logout", func(c *gin.Context) {
		userController.SignOut(c)
	})
	users.POST("", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		userController.Create(c, b)
	})

	// Articles
	articles := v1.Group("/articles")
	articles.Use(middleware.ResourcePermissionMiddleware())
	articleController := controllers.NewArticleController(sqlhandler)
	articles.POST("", func(c *gin.Context) {
		b := binding.Default(c.Request.Method, c.ContentType())
		articleController.Create(c, b)
	})
	articles.GET("", func(c *gin.Context) {
		articleController.Index(c)
	})
	articles.GET("/:id", func(c *gin.Context) {
		articleController.Show(c)
	})
	articles.DELETE("/:id", func(c *gin.Context) {
		articleController.Destroy(c)
	})
	Router = router
}
