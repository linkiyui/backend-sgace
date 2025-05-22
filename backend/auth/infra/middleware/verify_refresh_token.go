package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sgace/backend/auth/auth_token"
)

func VerifyRecoverToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	user_id, err := auth_token.ValidateRefreshToken(strings.TrimPrefix(token, "Bearer "))
	if err != nil {

		c.AbortWithStatus(401)
		return
	}
	c.Set("user_id", user_id)
	c.Next()
}
