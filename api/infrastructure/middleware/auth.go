package middleware

import (
	"shokubunka-media/api/domain"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandleFunc {
	return func(c *gin.Context) {
		session := sessions.Default()
		userID := session.Get("userID")
		if userID == nil || userID == 0 {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		setCurrentUser(c, userID)
	}
}

func setCurrentUser(c *gin.Context, userID) {
	var currentUser domain.User
	db := DB()
	db.First(&currentUser, userID)
	c.Set("current_user", currentUser)
}

func DB() *gorm.DB {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	db, _ := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	return db
}