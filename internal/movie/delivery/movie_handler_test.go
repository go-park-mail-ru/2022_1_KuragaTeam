package delivery

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/mock"
	"net/http/httptest"
	"testing"
)

func TestMovieDelivery_GetMainMovie(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := internal.MainMovieInfoDTO{
		ID:      0,
		Name:    "Movie1",
		Tagline: "This is test movie",
		Picture: "movie_picture.webp",
	}

	tests := []struct {
		name          string
		param         string
		useCaseMock   *mock.MockMovieService
		expected      internal.MainMovieInfoDTO
		expectedError bool
	}{
		{
			name:  "get main movie",
			param: "1",
			useCaseMock: &mock.MockMovieService{
				GetMainMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return &movie, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/mainMovie", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)

			r := NewHandler(test.useCaseMock)
			mainMovie := r.GetMainMovie()

			_ = mainMovie(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result internal.MainMovieInfoDTO
			_ = json.Unmarshal(body.Bytes(), &result)
			if test.expectedError {
				assert.Equal(t, "500 Internal Server Error", status)
			} else {
				assert.Equal(t, test.expected, result)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}
