package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yuudi/ero-runner/server/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Backed with Trusted Header by Authelia
		userID := c.GetHeader("Remote-User")
		if userID == "" {
			c.AbortWithStatus(401)
			return
		}

		name := c.GetHeader("Remote-Name")
		user, err := model.GetOrCreateUser(c, userID, name)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
