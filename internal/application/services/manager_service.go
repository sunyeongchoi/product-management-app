package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"product-management/internal/domain/entities"
	"product-management/internal/domain/repositories"
	"product-management/internal/interface/api/rest/request"
	"strconv"
	"time"
)

type ManagerService struct {
	repository repositories.ManagerRepository
}

func NewManagerService(repository repositories.ManagerRepository) *ManagerService {
	return &ManagerService{repository: repository}
}

func (m *ManagerService) SignUp(managerRequest *request.CreateManagerRequest) (statusCode int, err error) {
	// 휴대폰번호 중복체크
	_, err = m.repository.Get(managerRequest.Phone)
	if err == nil {
		return http.StatusBadRequest, errors.New("중복된 핸드폰번호입니다.")
	}
	var newManager = entities.NewManager(0, managerRequest.Phone, managerRequest.Password)
	validatedManager, err := entities.NewValidatedManager(newManager)
	if err != nil {
		return http.StatusBadRequest, err
	}
	// 비밀번호 암호화
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validatedManager.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadRequest, errors.New("비밀번호 암호화 오류")
	}
	validatedManager.Password = string(hashedPassword)

	err = m.repository.SignUp(validatedManager)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (m *ManagerService) Login(managerRequest *request.CreateManagerRequest) (token string, tokenExpiration time.Time, statusCode int, err error) {
	var newManager = entities.NewManager(0, managerRequest.Phone, managerRequest.Password)
	validatedManager, err := entities.NewValidatedManager(newManager)
	if err != nil {
		return "", time.Time{}, http.StatusBadRequest, err
	}

	managerFromDB, err := m.repository.Get(validatedManager.Phone)
	if err != nil {
		return "", time.Time{}, http.StatusUnauthorized, err
	}
	// 비밀번호 검증
	err = bcrypt.CompareHashAndPassword([]byte(managerFromDB.Password), []byte(validatedManager.Password))
	if err != nil {
		return "", time.Time{}, http.StatusUnauthorized, err
	}
	// JWT 생성
	token, tokenExpiration, err = CreateToken(managerFromDB.ID, managerFromDB.Phone)
	if err != nil {
		return "", time.Time{}, http.StatusUnauthorized, err
	}
	return token, tokenExpiration, http.StatusOK, nil
}

type Claims struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func CreateToken(id int, phone string) (token string, tokenExpiration time.Time, err error) {
	// TODO: Refresh Token 고려 필요
	now := time.Now()
	timeDurationStr := os.Getenv("JWT_TIME_DURATION")
	timeDuration, err := strconv.Atoi(timeDurationStr)
	if err != nil {
		return "", time.Time{}, errors.New("JWT_TIME_DURATION을 올바른 타입으로 설정해주세요.")
	}
	tokenExpiration = now.Add(time.Duration(timeDuration) * time.Second)
	claims := &Claims{
		ID:    id,
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiration),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", time.Time{}, err
	}
	return token, tokenExpiration, nil
}

func GetClaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}