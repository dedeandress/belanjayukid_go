package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckHash(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
