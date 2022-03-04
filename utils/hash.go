package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	}

	return true, nil
}
