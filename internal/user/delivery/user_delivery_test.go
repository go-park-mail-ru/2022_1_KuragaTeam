package delivery

//
//import (
//	"myapp/internal/user"
//	"myapp/internal/user/delivery/mock"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//func TestUserDelivery_SignUp(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//
//	mockService := mock.NewMockService(ctl)
//
//	tests := []struct {
//		name           string
//		mock           func()
//		expectedStatus int
//		expectedJSON   string
//	}{
//		{
//			name: "Handler returned status 201",
//			mock: func() {
//				userData := &user.CreateUserDTO{
//					Name:     "Olga",
//					Email:    "olga@mail.ru",
//					Password: "olga123321",
//				}
//				gomock.InOrder(
//					mockService.EXPECT().SignUp(userData).Return("session", "", nil),
//				)
//			},
//			expectedStatus: http.StatusCreated,
//			expectedJSON:   "{\"status\":201,\"message\":\"User created\"}\n",
//		},
//		//{
//		//	name: "Handler returned status 401, usecase.Login returned CustomError with ErrorType = 400",
//		//	usecaseMock: &mock.MockUserUseCase{
//		//		RegisterFunc: func(user models.User) (string, *models.CustomError) {
//		//			return "", &models.CustomError{
//		//				ErrorType: 400,
//		//				Message:   "BadRequest",
//		//			}
//		//		},
//		//	},
//		//	expectedStatus: http.StatusOK,
//		//	expectedJSON:   "{\"status\":400,\"message\":\"BadRequest\"}\n",
//		//},
//		//{
//		//	name: "Handler returned status 500, usecase.Login returned CustomError with ErrorType = 500",
//		//	usecaseMock: &mock.MockUserUseCase{
//		//		RegisterFunc: func(user models.User) (string, *models.CustomError) {
//		//			return "", &models.CustomError{
//		//				ErrorType:     500,
//		//				OriginalError: errors.New("error"),
//		//			}
//		//		},
//		//	},
//		//	expectedStatus: http.StatusInternalServerError,
//		//	expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
//		//},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.POST, "/api/v1/user/signup", strings.NewReader(`{"email": "test.inter@ndeiud.com", "password": "jfdIHD#&n873D"}`))
//			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//
//			th := test
//			th.mock()
//
//			handler := NewHandler(mockService, logger)
//			handler.Register(server)
//
//			if assert.NoError(t, r.Register(ctx)) {
//				assert.Equal(t, tt.expectedStatus, rec.Code)
//				assert.Equal(t, tt.expectedJSON, rec.Body.String())
//			}
//		})
//	}
//}
