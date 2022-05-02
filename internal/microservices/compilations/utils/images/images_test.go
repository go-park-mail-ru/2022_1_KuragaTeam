package images

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			result:   fmt.Sprintf("http://%s/%s/%s", os.Getenv("HOST")+"/api/v1/minio", "bucket1", "name1"),
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
