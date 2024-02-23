package manager

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

const jwtKey = "example"

func CreateToken(id int, phone string) (token string, err error) {
	now := time.Now()
	// TODO: time expiration 입력받도록 변경
	tokenExpiration := now.Add(time.Duration(50000) * time.Second)
	claims := &Claims{
		ID:    id,
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiration),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetClaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
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