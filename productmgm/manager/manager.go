package manager

import (
	"errors"
	"net/http"
	"regexp"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"product-management/productmgm"
	"product-management/server"
	"product-management/server/manager"
)

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

func SignUp(mng manager.Manager) (statusCode int, err error) {
	// 휴대폰번호 중복체크
	_, err = getManagerDBConn().Get(mng.Phone)
	if err == nil {
		return http.StatusBadRequest, errors.New("중복된 핸드폰번호입니다.")
	}
	// 휴대폰번호 정규식 체크
	if !isValidPhoneNumber(mng.Phone) {
		return http.StatusBadRequest, errors.New("올바른 휴대폰번호 형식이 아닙니다. 01012345678 형태로 입력해주세요.")
	}
	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mng.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadRequest, err
	}
	mng.Password = string(hashedPassword)

	err = getManagerDBConn().SignUp(mng)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
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

func Login(mng manager.Manager) (token string, tokenExpiration time.Time, statusCode int, err error) {
	managerFromDB, err := getManagerDBConn().Get(mng.Phone)
	if err != nil {
		return "", time.Time{}, http.StatusInternalServerError, err
	}
	// 비밀번호 검증
	err = bcrypt.CompareHashAndPassword([]byte(managerFromDB.Password), []byte(mng.Password))
	if err != nil {
		return "", time.Time{}, http.StatusInternalServerError, err
	}
	// JWT 생성
	token, tokenExpiration, err = productmgm.CreateToken(managerFromDB.ID, managerFromDB.Phone)
	if err != nil {
		return "", time.Time{}, http.StatusInternalServerError, err
	}
	return token, tokenExpiration, http.StatusOK, nil
}