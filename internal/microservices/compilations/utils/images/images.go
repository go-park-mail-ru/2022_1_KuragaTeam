package images

import (
	"fmt"
	"os"
)

func GenerateFileURL(fileName string, bucket string) (string, error) {
	return fmt.Sprintf("http://%s/api/v1/%s/%s", os.Getenv("NGINX"), bucket, fileName), nil
}
