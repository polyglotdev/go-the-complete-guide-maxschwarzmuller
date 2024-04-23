package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/utils"
)

func CheckAuth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	// checkov:skip=CKV_SECRET_6: not a password

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		return
	}

	c.Set("userId", userId)
	c.Next()

}
