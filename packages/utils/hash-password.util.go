package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), error
}

func VerifyPassword(password string, hashPhrase []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPhrase, []byte(password))
	return err == nil
}
