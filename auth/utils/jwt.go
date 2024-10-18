package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	expTime int64             = 30 * 24 * 60 * 60
	method  jwt.SigningMethod = jwt.SigningMethodHS256
)

func Issue(email string) (tokenString string, err error) {
	exp := time.Now().Unix() + expTime

	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"exp": exp,
		"sub": email,
	})

	tokenString, err = token.SignedString(os.Getenv("JWT_SECRET"))

	return tokenString, err
}
