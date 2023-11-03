package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenHash(password string) (string, error) {
	bp, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bp), nil
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}