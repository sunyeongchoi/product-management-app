package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/cmd/manager"
)

func TokenAuthMiddleware(c *gin.Context)  {
	jwtTokenCookie, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	claims, err := manager.GetClaims(jwtTokenCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("ID", claims.ID)
	c.Set("PHONE", claims.Phone)
}