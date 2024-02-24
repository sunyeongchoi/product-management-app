package manager

import (
	"net/http"
	"sync"

	"product-management/common"

	"product-management/models"
	"product-management/sql"
	managers "product-management/sql/manager"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type apiManager struct{}

func GetManagerAPIManager() *apiManager {
	return &apiManager{}
}

var (
	once          sync.Once
	managerDBConn *managers.DBManagerService
)

func getManagerDBConn() *managers.DBManagerService {
	once.Do(func() {
		managerDBConn = managers.NewDBManagerService(sql.DBConn)
	})
	return managerDBConn
}

func (m apiManager) SignUp(c *gin.Context) {
	var manager models.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	// 휴대폰번호 중복체크
	_, err := getManagerDBConn().Get(manager.Phone)
	if err == nil {
		common.NewManagerResponse(http.StatusInternalServerError, "중복된 핸드폰번호입니다.", nil).GetManagerResponse(c)
		return
	}
	// TODO: 휴대폰번호 정규식 체크
	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(manager.Password), bcrypt.DefaultCost)
	if err != nil {
		common.NewManagerResponse(http.StatusInternalServerError, err.Error(), nil).GetManagerResponse(c)
		return
	}
	manager.Password = string(hashedPassword)

	err = getManagerDBConn().SignUp(manager)
	if err != nil {
		common.NewManagerResponse(http.StatusInternalServerError, err.Error(), nil).GetManagerResponse(c)
		return
	}
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) Login(c *gin.Context) {
	var managerFromInput models.Manager
	if err := c.ShouldBindJSON(&managerFromInput); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	managerFromDB, err := getManagerDBConn().Get(managerFromInput.Phone)
	if err != nil {
		common.NewManagerResponse(http.StatusInternalServerError, err.Error(), nil).GetManagerResponse(c)
		return
	}
	// 비밀번호 검증
	err = bcrypt.CompareHashAndPassword([]byte(managerFromDB.Password), []byte(managerFromInput.Password))
	if err != nil {
		common.NewManagerResponse(http.StatusInternalServerError, err.Error(), nil).GetManagerResponse(c)
		return
	}
	// JWT 생성
	token, err := CreateToken(managerFromDB.ID, managerFromDB.Phone)
	if err != nil {
		return
	}
	c.SetCookie(common.JWTTOKEN, token, 50000, "/", "localhost:8080", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) LogOut(c *gin.Context) {
	c.SetCookie(common.JWTTOKEN, "", -1, "/", "localhost:8080", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}
