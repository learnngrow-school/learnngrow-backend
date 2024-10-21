package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ExpTime int64             = 30 * 24 * 60 * 60
	method  jwt.SigningMethod = jwt.SigningMethodHS256
)

func Issue(email string) (tokenString string, err error) {
	exp := time.Now().Unix() + ExpTime

	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"exp": exp,
		"sub": email,
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}
