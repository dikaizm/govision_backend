package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(secretKey string, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+secretKey), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func CheckPasswordHash(secretKey string, password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+secretKey))
	if err != nil {
		return err
	}

	return nil
}
