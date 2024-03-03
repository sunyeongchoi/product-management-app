package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/internal/application/interfaces"
	"product-management/internal/interface/api/rest/request"
	"product-management/utils"
)

type ManagerController struct {
	service interfaces.ManagerService
}

func NewManagerController(c *gin.Engine, service interfaces.ManagerService) *ManagerController {
	controller := &ManagerController{
		service: service,
	}
	c.POST("/signup", controller.SignUp)
	c.POST("/login", controller.Login)
	c.POST("/logout", controller.LogOut)
	return controller
}

func (mc *ManagerController) SignUp(c *gin.Context) {
	var createManagerRequest request.CreateManagerRequest
	if err := c.ShouldBindJSON(&createManagerRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	managerCommand, err := createManagerRequest.ToCreateManagerCommand()
	if err != nil {
		utils.NewManagerResponse(http.StatusBadRequest, err.Error(), nil).GetManagerResponse(c)
		return
	}
	statusCode, err := mc.service.SignUp(managerCommand)
	if err != nil {
		utils.NewManagerResponse(statusCode, err.Error(), nil).GetManagerResponse(c)
		return
	}
	utils.NewManagerResponse(http.StatusOK, utils.OKAYMSG, nil).GetManagerResponse(c)
}

func (mc *ManagerController) Login(c *gin.Context) {
	var createManagerRequest request.CreateManagerRequest
	if err := c.ShouldBindJSON(&createManagerRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	managerCommand, err := createManagerRequest.ToCreateManagerCommand()
	if err != nil {
		utils.NewManagerResponse(http.StatusBadRequest, err.Error(), nil).GetManagerResponse(c)
		return
	}
	fmt.Println("managerCommand", managerCommand)
	token, tokenExpiration, statusCode, err := mc.service.Login(managerCommand)
	if err != nil {
		utils.NewManagerResponse(statusCode, err.Error(), nil).GetManagerResponse(c)
		return
	}
	c.SetCookie(utils.JWTTOKEN, token, int(tokenExpiration.Unix()), "/", "", false, true)
	utils.NewManagerResponse(http.StatusOK, utils.OKAYMSG, nil).GetManagerResponse(c)
}

func (m *ManagerController) LogOut(c *gin.Context) {
	c.SetCookie(utils.JWTTOKEN, "", -1, "/", "", false, true)
	utils.NewManagerResponse(http.StatusOK, utils.OKAYMSG, nil).GetManagerResponse(c)
}