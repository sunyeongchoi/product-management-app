package middleware

import (
	"net/http"

	"product-management/pkg/apiclient/manager"
	"product-management/productmgm/common"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	jwtTokenCookie, err := c.Cookie(common.JWTTOKEN)
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