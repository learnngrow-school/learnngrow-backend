package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ExpTime int64             = 30 * 24 * 60 * 60
	method  jwt.SigningMethod = jwt.SigningMethodHS256
	secret  []byte = []byte(os.Getenv("JWT_SECRET"))
)

func Issue(email string) (tokenString string, err error) {
	exp := time.Now().Unix() + ExpTime

	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"exp": exp,
		"sub": email,
	})

	tokenString, err = token.SignedString(secret)

	return tokenString, err
}

func GetData(tokenString string) (email jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method is bad")
		}
		return secret, nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token invalid somehow")
	}

	return claims, nil
}
