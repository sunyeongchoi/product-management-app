package middleware

import (
	"net/http"
	"product-management/internal/application/services"
	"product-management/utils"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	jwtTokenCookie, err := c.Cookie(utils.JWTTOKEN)
	if err != nil {
		utils.NewProductResponse(http.StatusUnauthorized, err.Error(), nil).GetProductResponse(c)
		c.Abort()
		return
	}
	_, err = services.GetClaims(jwtTokenCookie)
	if err != nil {
		utils.NewProductResponse(http.StatusUnauthorized, err.Error(), nil).GetProductResponse(c)
		c.Abort()
		return
	}
}