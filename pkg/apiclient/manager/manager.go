package manager

import (
	"net/http"

	"product-management/productmgm/common"
	servicemanager "product-management/productmgm/manager"
	"product-management/server/manager"

	"github.com/gin-gonic/gin"
)

type apiManager struct{}

func GetManagerAPIManager() *apiManager {
	return &apiManager{}
}

func (m apiManager) SignUp(c *gin.Context) {
	var mng manager.Manager
	if err := c.ShouldBindJSON(&mng); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	statusCode, err := servicemanager.SignUp(mng)
	if err != nil {
		common.NewManagerResponse(statusCode, err.Error(), nil).GetManagerResponse(c)
		return
	}
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) Login(c *gin.Context) {
	var managerFromInput manager.Manager
	if err := c.ShouldBindJSON(&managerFromInput); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	token, tokenExpiration, statusCode, err := servicemanager.Login(managerFromInput)
	if err != nil {
		common.NewManagerResponse(statusCode, err.Error(), nil).GetManagerResponse(c)
		return
	}
	c.SetCookie(common.JWTTOKEN, token, int(tokenExpiration.Unix()), "/", "", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) LogOut(c *gin.Context) {
	c.SetCookie(common.JWTTOKEN, "", -1, "/", "", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}
