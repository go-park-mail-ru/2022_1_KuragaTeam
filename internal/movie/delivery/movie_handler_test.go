package delivery

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"myapp/internal"
	mock_movie "myapp/internal/movie/mock"
	"net/http"
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

	useCaseMock := mock_movie.NewMockService(ctrl)

	movie := internal.MainMovieInfoDTO{
		ID:      0,
		Name:    "Movie1",
		Tagline: "This is test movie",
		Picture: "movie_picture.webp",
	}
	//tests := []struct {
	//	name          string
	//	param         string
	//	useCaseMock   *mock_movie.MockService
	//	expected      internal.MainMovieInfoDTO
	//	expectedError bool
	//}{
	//	{
	//		name:          "get home",
	//		param:         "1",
	//		useCaseMock:   &mock_movie.MockService{ctrl},
	//		expected:      movie,
	//		expectedError: false,
	//	},
	//	{
	//		name:          "GetHome() error",
	//		param:         "1",
	//		useCaseMock:   &mock_movie.MockService{},
	//		expected:      movie,
	//		expectedError: true,
	//	},
	//}

	//for _, test := range tests {
	t.Run("Test", func(t *testing.T) {
		server := echo.New()

		useCaseMock.EXPECT().GetMainMovie().DoAndReturn(func() (*internal.MainMovieInfoDTO, error) {
			return &movie, nil
		})
		delivery := NewHandler(useCaseMock)

		response := httptest.NewRequest(echo.GET, "/api/v1/mainMovie", nil)
		rec := httptest.NewRecorder()
		ctx := server.NewContext(response, rec)

		delivery.Register(server)
		handlerFunc := delivery.GetMainMovie()
		err := handlerFunc(ctx)
		assert.NoError(t, err)

		assert.Equal(t, rec.Code, http.StatusOK)
		body, _ := ioutil.ReadAll(rec.Result().Body)
		var receive internal.MainMovieInfoDTO
		err = json.Unmarshal(body, &receive)
		assert.NoError(t, err)
		assert.Equal(t, movie, receive)

		//_ = delivery.Home(ctx)
		//body := recorder.Body
		//status := recorder.Result().Status
		//var result []models.Track
		//_ = json.Unmarshal(body.Bytes(), &result)
		//if test.expectedError {
		//	assert.Equal(t, "500 Internal Server Error", status)
		//} else {
		//	assert.Equal(t, test.expected, result)
		//	assert.Equal(t, "200 OK", status)
		//}
	})
	//}
}
