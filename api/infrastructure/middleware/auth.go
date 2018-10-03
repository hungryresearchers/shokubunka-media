package middleware

import (
	"api/domain"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		id, ok := userID.(int)
		if !ok {
			return
		}
		setCurrentUser(c, id)
	}
}

func setCurrentUser(c *gin.Context, userID int) {
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
