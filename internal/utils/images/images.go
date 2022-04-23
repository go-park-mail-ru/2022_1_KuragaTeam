package images

import (
	"fmt"
	"myapp/internal/models"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GenerateObjectName(input models.UploadInput) string {
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

	return fmt.Sprintf("http://%s/api/v1/%s/%s", os.Getenv("NGINX"), bucket, fileName), nil
}
