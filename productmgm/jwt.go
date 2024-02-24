package productmgm

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
