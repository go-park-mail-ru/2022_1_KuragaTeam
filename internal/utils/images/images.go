package images

import (
	"fmt"
	"myapp/internal/user"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GenerateObjectName(input user.UploadInput) string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(input.UserID)), formatted, "webp")
}

func GenerateFileURL(fileName string, bucket string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s/%s/%s", os.Getenv("MINIOURL"), bucket, fileName), nil
}