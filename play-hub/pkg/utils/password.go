package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GetHashedPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password []byte, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), password)
	return err == nil
}

//func IsValidPassword(password string) bool {
//	var hasUpper, hasLower, hasSpecial bool
//
//	if len(password) < 8 {
//		return false
//	}
//
//	for _, char := range password {
//		switch {
//		case unicode.IsUpper(char):
//			hasUpper = true
//		case unicode.IsLower(char):
//			hasLower = true
//		case unicode.IsPunct(char) || unicode.IsSymbol(char):
//			hasSpecial = true
//		}
//	}
//
//	return hasUpper && hasLower && hasSpecial
//}
