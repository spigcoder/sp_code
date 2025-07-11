package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(encryptText string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(encryptText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashStr), err
}

func CompareHashAndPassword(hashPassword, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password))

	return err == nil
}
