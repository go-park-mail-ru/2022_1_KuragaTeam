package images

import (
	"fmt"
	"os"
)

func GenerateFileURL(fileName string, bucket string) (string, error) {
	return fmt.Sprintf("http://%s/api/v1/minio/%s/%s", os.Getenv("HOST"), bucket, fileName), nil
}
