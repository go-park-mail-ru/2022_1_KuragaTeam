package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/mock"
	"testing"
)

func TestMovieUsecase_GetMainMovie(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
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
		storageMock   *mock.MockMovieStorage
		expected      *internal.MainMovieInfoDTO
		expectedError bool
	}{
		{
			name: "Get main movie",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return &movie, nil
				},
			},
			expected:      &movie,
			expectedError: false,
		},
		{
			name: "Return error",
			storageMock: &mock.MockMovieStorage{
				GetRandomMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.storageMock, nil, nil, nil)
			mainMovie, err := r.GetMainMovie()

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, mainMovie)
			}
		})
	}
}
