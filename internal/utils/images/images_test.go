package images

import (
	"fmt"
	"myapp/internal/user"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateObjectName(t *testing.T) {
	tests := []struct {
		name   string
		user   user.UploadInput
		result string
	}{
		{
			name: "GenerateObjectName",
			user: user.UploadInput{
				UserID:      0,
				File:        nil,
				Size:        0,
				ContentType: "",
			},
			result: fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(0)),
				fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
					time.Now().Minute(), time.Now().Second()),
				"webp"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			result := GenerateObjectName(th.user)

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
			result:   fmt.Sprintf("http://%s/%s/%s", "localhost:8000/api/v1", "bucket1", "name1"),
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
