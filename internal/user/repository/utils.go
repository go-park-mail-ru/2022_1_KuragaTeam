package repository

import (
	"fmt"
	"myapp/constants"
	"myapp/internal/user"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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

func GenerateObjectName(input user.UploadInput) string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(input.UserID)), formatted, "webp")
}

func GenerateFileURL(fileName string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", os.Getenv("MINIOURL"), constants.UserObjectsBucketName, fileName), nil
}
