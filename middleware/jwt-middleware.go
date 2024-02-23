package middleware

import (
	"net/http"

	"product-management/cmd/manager"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	jwtTokenCookie, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	_, err = manager.GetClaims(jwtTokenCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
}