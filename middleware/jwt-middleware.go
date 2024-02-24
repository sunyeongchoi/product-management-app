package middleware

import (
	"net/http"

	"product-management/common"

	"product-management/cmd/manager"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	jwtTokenCookie, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		common.NewProductResponse(http.StatusUnauthorized, err.Error(), nil).GetProductResponse(c)
		c.Abort()
		return
	}
	_, err = manager.GetClaims(jwtTokenCookie)
	if err != nil {
		common.NewProductResponse(http.StatusUnauthorized, err.Error(), nil).GetProductResponse(c)
		c.Abort()
		return
	}
}