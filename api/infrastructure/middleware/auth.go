package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandleFunc {
	return func(c *gin.Context) {
		session := sessions.Default()
	}
}
