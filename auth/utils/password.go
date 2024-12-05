package auth

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password []byte) ([]byte, error) {
	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	return bcrypt.GenerateFromPassword(password, cost)
}

func Check(given string, actual []byte) error {
	return bcrypt.CompareHashAndPassword(actual, []byte(given))
}
