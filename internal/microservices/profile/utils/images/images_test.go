package images

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateObjectName(t *testing.T) {
	tests := []struct {
		name   string
		userID int64
		result string
	}{
		{
			name:   "GenerateObjectName",
			userID: int64(0),
			result: fmt.Sprintf("%s_%s.%s", strconv.Itoa(0),
				fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
					time.Now().Minute(), time.Now().Second()),
				"webp"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			result := GenerateObjectName(th.userID)

			assert.Equal(t, th.result, result)
		})
	}
}

func TestGenerateFileURL(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		bucket   string
		result   string
	}{
		{
			name:     "GenerateFileUR",
			fileName: "name1",
			bucket:   "bucket1",
			result:   fmt.Sprintf("https://%s/%s/%s", os.Getenv("HOST")+"/api/v1/minio", "bucket1", "name1"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			result, err := GenerateFileURL(th.fileName, th.bucket)

			assert.NoError(t, err)
			assert.Equal(t, th.result, result)
		})
	}
}
