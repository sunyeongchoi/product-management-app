package manager

import (
	"net/http"
	"regexp"
	"sync"

	"product-management/productmgm/common"
	"product-management/server"
	"product-management/server/manager"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type apiManager struct{}

func GetManagerAPIManager() *apiManager {
	return &apiManager{}
}

var (
	once          sync.Once
	managerDBConn *manager.DBManagerService
)

func getManagerDBConn() *manager.DBManagerService {
	once.Do(func() {
		managerDBConn = manager.NewDBManagerService(server.DBConn)
	})
	return managerDBConn
}

func (m apiManager) SignUp(c *gin.Context) {
	var mng manager.Manager
	if err := c.ShouldBindJSON(&mng); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	// 휴대폰번호 중복체크
	_, err := getManagerDBConn().Get(mng.Phone)
	if err == nil {
		common.NewManagerResponse(http.StatusBadRequest, "중복된 핸드폰번호입니다.", nil).GetManagerResponse(c)
		return
	}
	// 휴대폰번호 정규식 체크
	if !isValidPhoneNumber(mng.Phone) {
		common.NewManagerResponse(http.StatusBadRequest, "올바른 휴대폰번호 형식이 아닙니다. 01012345678 형태로 입력해주세요.", nil).GetManagerResponse(c)
		return
	}
	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mng.Password), bcrypt.DefaultCost)
	if err != nil {
		common.NewManagerResponse(http.StatusBadRequest, err.Error(), nil).GetManagerResponse(c)
		return
	}
	mng.Password = string(hashedPassword)

	err = getManagerDBConn().SignUp(mng)
	if err != nil {
		common.NewManagerResponse(http.StatusInternalServerError, err.Error(), nil).GetManagerResponse(c)
		return
	}
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) Login(c *gin.Context) {
	var managerFromInput manager.Manager
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
	token, tokenExpiration, err := CreateToken(managerFromDB.ID, managerFromDB.Phone)
	if err != nil {
		return
	}
	c.SetCookie(common.JWTTOKEN, token, int(tokenExpiration.Unix()), "/", "", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func (m apiManager) LogOut(c *gin.Context) {
	c.SetCookie(common.JWTTOKEN, "", -1, "/", "", false, true)
	common.NewManagerResponse(http.StatusOK, common.OKAYMSG, nil).GetManagerResponse(c)
}

func isValidPhoneNumber(phoneNumber string) bool {
	//^: 문자열의 시작
	//01: 휴대폰 번호가 010으로 시작
	//[0-9]{8,9}: 8 또는 9자리의 숫자 그룹
	//$: 문자열의 끝
	regexPattern := `^01[0-9]{8,9}$`
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(phoneNumber)
}
