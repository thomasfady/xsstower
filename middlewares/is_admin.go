package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAdminMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		session := sessions.Default(g)
		is_admin := session.Get("is_admin")
		if is_admin == nil || is_admin != true {
			g.AbortWithStatus(403)
			return
		}
		g.Next()
	}
}
