package delivery

import (
	"bytes"
	"errors"
	"mime/multipart"
	"myapp/internal/csrf"
	"myapp/internal/user"
	"myapp/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestUserDelivery_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		data           string
	}{
		{
			name: "Handler returned status 201",
			mock: func() {
				userData := &user.CreateUserDTO{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(userData).Return("session", "", nil),
				)
			},
			expectedStatus: http.StatusCreated,
			expectedJSON:   "{\"status\":201,\"message\":\"User created\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 400, Email is not unique",
			mock: func() {
				userData := &user.CreateUserDTO{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(userData).Return("", "ERROR: Email is not unique", nil),
				)
			},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"ERROR: Email is not unique\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 500, usecase SignUp error",
			mock: func() {
				userData := &user.CreateUserDTO{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(userData).Return("", "", errors.New("error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/signup", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			signup := handler.SignUp()
			_ = signup(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_LogIn(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		data           string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &user.LogInUserDTO{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(userData).Return("session", nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"User can be logged in\"}\n",
			expectedError:  false,
			data:           `{"email": "olga@mail.ru", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 404, User not found",
			mock: func() {
				userData := &user.LogInUserDTO{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(userData).Return("", nil),
				)
			},
			expectedStatus: http.StatusNotFound,
			expectedJSON:   "{\"status\":404,\"message\":\"User not found\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 500, usecase LogIn error",
			mock: func() {
				userData := &user.LogInUserDTO{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(userData).Return("", errors.New("error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/login", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			login := handler.LogIn()
			_ = login(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_GetUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		userIDKey      string
		userIDValue    int64
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &user.ProfileUserDTO{
					Name:   "Olga",
					Email:  "olga@mail.ru",
					Avatar: "avatar",
				}
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(int64(1)).Return(userData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"user\":{\"username\":\"Olga\",\"email\":\"olga@mail.ru\",\"avatar\":\"avatar\"}}\n",
			expectedError:  false,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			expectedError:  true,
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
		},
		{
			name: "Handler returned status 500, usecase GetUserProfile error",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(int64(1)).Return(nil, errors.New("error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/profile", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			profile := handler.GetUserProfile()
			_ = profile(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_LogOut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		cookie         http.Cookie
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().LogOut("session").Return(nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"User is logged out\"}\n",
			expectedError:  false,
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
		},
		{
			name:           "Handler returned status 500",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
			cookie: http.Cookie{
				Name:     "Wrong_Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
		},
		{
			name: "Handler returned status 500, usecase LogOut error",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().LogOut("session").Return(errors.New("error")),
				)
			},
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			expectedError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.DELETE, "/api/v1/logout", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")

			ctx.Request().AddCookie(&th.cookie)

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			logout := handler.LogOut()
			_ = logout(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_EditAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		userIDKey      string
		userIDValue    int64
		fieldName      string
		content        []byte
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			expectedError:  true,
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
		},
		{
			name:           "Handler returned status 500, Error file",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			fieldName:      "wrong_file",
		},
		{
			name:           "Handler returned status 500, Error content type",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			fieldName:      "file",
			content:        []byte("content"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			body_file := new(bytes.Buffer)
			writer := multipart.NewWriter(body_file)
			part, _ := writer.CreateFormFile(th.fieldName, "avatar_test.webp")
			part.Write(th.content)
			writer.Close()

			req := httptest.NewRequest(echo.PUT, "/api/v1/avatar", body_file)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			avatar := handler.EditAvatar()
			_ = avatar(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_EditProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		data           string
		userIDKey      string
		userIDValue    int64
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			expectedError:  true,
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"username": "Olga", "password": "olga123321"}`,
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			data:           `{"username": "Olga", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &user.EditProfileDTO{
					ID:       int64(1),
					Name:     "Olga",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().EditProfile(userData).Return(nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"Profile is edited\"}\n",
			expectedError:  false,
			data:           `{"username": "Olga", "password": "olga123321"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
		{
			name: "Handler returned status 500, usecase EditProfile error",
			mock: func() {
				userData := &user.EditProfileDTO{
					ID:       int64(1),
					Name:     "Olga",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().EditProfile(userData).Return(errors.New("error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
			data:           `{"username": "Olga", "password": "olga123321"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.PUT, "/api/v1/edit", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			edit := handler.EditProfile()
			_ = edit(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_GetCsrf(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	token, _ := csrf.Tokens.Create("session", time.Now().Add(time.Hour).Unix())

	tests := []struct {
		name           string
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		cookie         http.Cookie
	}{
		{
			name:           "Handler returned status 200",
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"" + token + "\"}\n",
			expectedError:  false,
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
		},
		{
			name:           "Handler returned status 500",
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
			cookie: http.Cookie{
				Name:     "Wrong_Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/csrf", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			ctx.Request().AddCookie(&th.cookie)

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			funcCSRF := handler.GetCsrf()
			_ = funcCSRF(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}

func TestUserDelivery_Auth(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		userIDKey      string
		userIDValue    int64
	}{
		//{
		//	name: "Handler returned status 200",
		//	mock: func() {
		//		gomock.InOrder(
		//			mockService.EXPECT().GetAvatar(int64(1)).Return(constants.DefaultImage, nil),
		//		)
		//	},
		//	expectedStatus: http.StatusOK,
		//	expectedJSON:   "{\"status\":200,\"message\":\"ok\"}\n",
		//	expectedError:  false,
		//	userIDKey:      "USER_ID",
		//	userIDValue:    int64(1),
		//},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			expectedError:  true,
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
		},
		{
			name: "Handler returned status 500, usecase GetAvatar error",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().GetAvatar(int64(1)).Return("", errors.New("error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			expectedError:  true,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
		{
			name: "Handler returned status 403",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().GetAvatar(int64(1)).Return("wrong_avatar.webp", nil),
				)
			},
			expectedStatus: http.StatusForbidden,
			expectedJSON:   "{\"status\":403,\"message\":\"wrong avatar\"}\n",
			expectedError:  false,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/auth", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set("Req", "/avatars/default_avatar.webp")
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewHandler(mockService, logger)
			handler.Register(server)
			auth := handler.Auth()
			_ = auth(ctx)

			body := rec.Body.String()
			status := rec.Code

			if test.expectedError {
				assert.Equal(t, th.expectedStatus, status)
			} else {
				assert.Equal(t, test.expectedJSON, body)
				assert.Equal(t, th.expectedStatus, status)
			}
		})
	}
}
