package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPwd string, salt string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd+salt))
	if err != nil {
		return false, err
	}

	return true, nil
}
