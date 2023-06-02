package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		session := sessions.Default(g)
		if session.Get("user_id") == nil {
			g.Status(401)
			g.Abort()
			return
		} else {
			g.Set("user_id", int(session.Get("user_id").(uint)))
			g.Next()
		}

	}
}
