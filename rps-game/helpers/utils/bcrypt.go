package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptPassword(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

func CompareBcryptPassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	res := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if res != nil {
		return false
	}
	return true
}
