package delivery

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"myapp/internal/constants"
	"myapp/internal/csrf"
	"myapp/internal/microservices/profile/proto"
	"myapp/internal/microservices/profile/usecase"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestProfileHandler_GetUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.ProfileData{
					Name:   "Olga",
					Email:  "olga@mail.ru",
					Avatar: "avatar",
					Date:   "2022-05-18T22:26:17.289395Z",
				}
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(gomock.Any(), inputData).Return(userData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"user\":{\"username\":\"Olga\",\"email\":\"olga@mail.ru\",\"avatar\":\"avatar\",\"date\":\"2022-05-18T22:26:17.289395Z\"}}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name: "Handler returned status 500, bad connection",
			mock: func() {
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(gomock.Any(), inputData).Return(nil, status.Error(codes.Unavailable, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase GetUserProfile error",
			mock: func() {
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(gomock.Any(), inputData).Return(nil, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
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
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			profile := handler.GetUserProfile()
			_ = profile(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_EditAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		fieldName      string
		content        []byte
		requestID      string
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, Error file",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"http: no such file\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			fieldName:      "wrong_file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, Error content type",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"File type is not supported\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			fieldName:      "file",
			content:        []byte("content"),
			requestID:      "REQUEST_ID",
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
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			avatar := handler.EditAvatar()
			_ = avatar(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_EditProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		data           string
		userIDKey      string
		userIDValue    int64
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"username": "Olga", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			data:           `{"username": "Olga", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.EditProfileData{
					ID:       int64(1),
					Name:     "Olga",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().EditProfile(gomock.Any(), userData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"Profile is edited\"}",
			data:           `{"username": "Olga", "password": "olga123321"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase EditProfile error",
			mock: func() {
				userData := &proto.EditProfileData{
					ID:       int64(1),
					Name:     "Olga",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().EditProfile(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"username": "Olga", "password": "olga123321"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
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
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			edit := handler.EditProfile()
			_ = edit(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_GetCsrf(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	token, _ := csrf.Tokens.Create("session", time.Now().Add(time.Hour).Unix())

	tests := []struct {
		name           string
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		cookie         http.Cookie
		requestID      string
	}{
		{
			name:           "Handler returned status 200",
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"" + token + "\"}",
			expectedError:  false,
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			requestID: "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}",
			expectedError:  false,
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			requestID: "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 500",
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"http: named cookie not present\"}",
			expectedError:  true,
			cookie: http.Cookie{
				Name:     "Wrong_Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			requestID: "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/csrf", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Request().AddCookie(&th.cookie)

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			funcCSRF := handler.GetCsrf()
			_ = funcCSRF(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_Auth(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				data := &proto.UserID{ID: 1}
				returnData := &proto.FileName{Name: constants.DefaultImage}
				gomock.InOrder(
					mockService.EXPECT().GetAvatar(gomock.Any(), data).Return(returnData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"ok\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}",
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase GetAvatar error",
			mock: func() {
				data := &proto.UserID{ID: 1}
				gomock.InOrder(
					mockService.EXPECT().GetAvatar(gomock.Any(), data).Return(nil, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"rpc error: code = Internal desc = error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 403",
			mock: func() {
				data := &proto.UserID{ID: 1}
				returnData := &proto.FileName{Name: "wrong_avatar.webp"}
				gomock.InOrder(
					mockService.EXPECT().GetAvatar(gomock.Any(), data).Return(returnData, nil),
				)
			},
			expectedStatus: http.StatusForbidden,
			expectedJSON:   "{\"status\":403,\"message\":\"wrong avatar\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/auth", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set("Req", "/api/v1/minio/avatars/default_avatar.webp")
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			auth := handler.Auth()
			_ = auth(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_AddLike(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		data           string
		userIDKey      string
		userIDValue    int64
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.LikeData{
					UserID:  1,
					MovieID: 3,
				}
				gomock.InOrder(
					mockService.EXPECT().AddLike(gomock.Any(), userData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"" + constants.LikeIsEdited + "\"}",
			data:           `{"id": "3"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase AddLike error",
			mock: func() {
				userData := &proto.LikeData{
					UserID:  1,
					MovieID: 3,
				}
				gomock.InOrder(
					mockService.EXPECT().AddLike(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"id": "3"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.PUT, "/api/v1/like", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			like := handler.AddLike()
			_ = like(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_RemoveLike(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		data           string
		userIDKey      string
		userIDValue    int64
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.LikeData{
					UserID:  1,
					MovieID: 3,
				}
				gomock.InOrder(
					mockService.EXPECT().RemoveLike(gomock.Any(), userData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"" + constants.LikeIsRemoved + "\"}",
			data:           `{"id": "3"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase AddLike error",
			mock: func() {
				userData := &proto.LikeData{
					UserID:  1,
					MovieID: 3,
				}
				gomock.InOrder(
					mockService.EXPECT().RemoveLike(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"id": "3"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.PUT, "/api/v1/dislike", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			dislike := handler.RemoveLike()
			_ = dislike(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_GetFavorites(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				likes := make([]int64, 0)
				likes = append(likes, 1, 2, 3)
				favorite := &proto.Favorites{Id: likes}
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetFavorites(gomock.Any(), inputData).Return(favorite, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"favorites\":{\"id\":[1,2,3]}}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name: "Handler returned status 500, bad connection",
			mock: func() {
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetFavorites(gomock.Any(), inputData).Return(nil, status.Error(codes.Unavailable, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase GetFavorites error",
			mock: func() {
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetFavorites(gomock.Any(), inputData).Return(nil, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/likes", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			likes := handler.GetFavorites()
			_ = likes(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_Check(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				data := &proto.UserID{ID: 1}
				gomock.InOrder(
					mockService.EXPECT().IsSubscription(gomock.Any(), data).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"ok\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}",
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase IsSubscription error",
			mock: func() {
				data := &proto.UserID{ID: 1}
				gomock.InOrder(
					mockService.EXPECT().IsSubscription(gomock.Any(), data).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 403",
			mock: func() {
				data := &proto.UserID{ID: 1}
				gomock.InOrder(
					mockService.EXPECT().IsSubscription(gomock.Any(), data).Return(&proto.Empty{}, status.Error(codes.PermissionDenied, constants.ErrNoSubscription.Error())),
				)
			},
			expectedStatus: http.StatusForbidden,
			expectedJSON:   "{\"status\":403,\"message\":\"no subscription\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/check", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			check := handler.Check()
			_ = check(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_GetPaymentsToken(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		userIDKey      string
		userIDValue    int64
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				data := &proto.UserID{ID: 1}
				outputData := &proto.Token{Token: "token"}
				gomock.InOrder(
					mockService.EXPECT().GetPaymentsToken(gomock.Any(), data).Return(outputData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"token\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}",
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase GetPaymentsToken error",
			mock: func() {
				data := &proto.UserID{ID: 1}
				outputData := &proto.Token{}
				gomock.InOrder(
					mockService.EXPECT().GetPaymentsToken(gomock.Any(), data).Return(outputData, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"rpc error: code = Internal desc = error\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/payments/token", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			token := handler.GetPaymentsToken()
			_ = token(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_Payment(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		data           string
		userIDKey      string
		userIDValue    int64
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"token": "token"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			data:           `{"token": "token"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.CheckTokenData{
					Token: "token",
					Id:    1,
				}
				gomock.InOrder(
					mockService.EXPECT().CheckPaymentsToken(gomock.Any(), userData).Return(&proto.Empty{}, nil),
					mockService.EXPECT().CreatePayment(gomock.Any(), userData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"Payment is created\"}",
			data:           `{"token": "token"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase CheckPaymentsToken error",
			mock: func() {
				userData := &proto.CheckTokenData{
					Token: "token",
					Id:    1,
				}
				gomock.InOrder(
					mockService.EXPECT().CheckPaymentsToken(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"token": "token"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase CreatePayment error",
			mock: func() {
				userData := &proto.CheckTokenData{
					Token: "token",
					Id:    1,
				}
				gomock.InOrder(
					mockService.EXPECT().CheckPaymentsToken(gomock.Any(), userData).Return(&proto.Empty{}, nil),
					mockService.EXPECT().CreatePayment(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"token": "token"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 400, usecase CreatePayment wrong payment token",
			mock: func() {
				userData := &proto.CheckTokenData{
					Token: "token",
					Id:    1,
				}
				gomock.InOrder(
					mockService.EXPECT().CheckPaymentsToken(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.InvalidArgument, constants.ErrWrongToken.Error())),
				)
			},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"wrong payment token\"}",
			data:           `{"token": "token"}`,
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/payment", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			payment := handler.Payment()
			_ = payment(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_Subscribe(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	data := url.Values{}
	data.Set("label", "token")
	data.Set("withdraw_amount", "2")

	wrongData := url.Values{}
	wrongData.Set("label", "token")
	wrongData.Set("withdraw_amount", "a")

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		contentType    string
		input          url.Values
	}{
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			requestID:      "WRONG_REQUEST_ID",
			contentType:    echo.MIMEApplicationForm,
			input:          data,
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.Token{
					Token: "token",
				}

				subscribeData := &proto.SubscribeData{
					Token:  "token",
					Amount: float32(2),
				}
				gomock.InOrder(
					mockService.EXPECT().CheckToken(gomock.Any(), userData).Return(&proto.Empty{}, nil),
					mockService.EXPECT().CreateSubscribe(gomock.Any(), subscribeData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"Payment is created\"}",
			requestID:      "REQUEST_ID",
			contentType:    echo.MIMEApplicationForm,
			input:          data,
		},
		{
			name: "Handler returned status 415 wrong content type",
			mock: func() {
			},
			expectedStatus: http.StatusUnsupportedMediaType,
			expectedJSON:   "{\"status\":415,\"message\":\"Unsupported media type\"}",
			requestID:      "REQUEST_ID",
			contentType:    echo.MIMEApplicationJSON,
			input:          data,
		},
		{
			name: "Handler returned status 500 error ParseFloat",
			mock: func() {
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"strconv.ParseFloat: parsing \\\"a\\\": invalid syntax\"}",
			requestID:      "REQUEST_ID",
			contentType:    echo.MIMEApplicationForm,
			input:          wrongData,
		},
		{
			name: "Handler returned status 500, CheckToken fail",
			mock: func() {
				userData := &proto.Token{
					Token: "token",
				}
				gomock.InOrder(
					mockService.EXPECT().CheckToken(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			requestID:      "REQUEST_ID",
			contentType:    echo.MIMEApplicationForm,
			input:          data,
		},
		{
			name: "Handler returned status 500, CreateSubscribe fail",
			mock: func() {
				userData := &proto.Token{
					Token: "token",
				}

				subscribeData := &proto.SubscribeData{
					Token:  "token",
					Amount: float32(2),
				}
				gomock.InOrder(
					mockService.EXPECT().CheckToken(gomock.Any(), userData).Return(&proto.Empty{}, nil),
					mockService.EXPECT().CreateSubscribe(gomock.Any(), subscribeData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			requestID:      "REQUEST_ID",
			contentType:    echo.MIMEApplicationForm,
			input:          data,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/subscribe", strings.NewReader(th.input.Encode()))
			req.Header.Set(echo.HeaderContentType, th.contentType)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			subscribe := handler.Subscribe()
			_ = subscribe(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestProfileHandler_GetRating(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockProfileClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		requestID      string
		userIDKey      string
		param          string
		userIDValue    int64
		paramExists    bool
	}{
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"Session required\"}",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			param:          "1",
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			param:          "1",
			requestID:      "WRONG_REQUEST_ID",
			paramExists:    true,
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			param:          "1",
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
		{
			name: "Handler returned status 200",
			mock: func() {
				data := &proto.MovieRating{UserID: 1, MovieID: int64(1)}
				rating := &proto.Rating{Rating: 5}
				gomock.InOrder(
					mockService.EXPECT().GetMovieRating(gomock.Any(), data).Return(rating, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"rating\":5}",
			param:          "1",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
		{
			name: "Handler returned status 500, No MovieID in context",
			mock: func() {
			},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"No MovieID in context\"}",
			param:          "1.6",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
		{
			name: "Handler returned status 500, GetMovieRating fail",
			mock: func() {
				data := &proto.MovieRating{UserID: 1, MovieID: int64(1)}
				rating := &proto.Rating{Rating: 5}
				gomock.InOrder(
					mockService.EXPECT().GetMovieRating(gomock.Any(), data).Return(rating, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			param:          "1",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
		{
			name: "Handler returned status 404, rating nill",
			mock: func() {
				data := &proto.MovieRating{UserID: 1, MovieID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetMovieRating(gomock.Any(), data).Return(nil, nil),
				)
			},
			expectedStatus: http.StatusNotFound,
			expectedJSON:   "{\"status\":404,\"rating\":0}",
			param:          "1",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
			paramExists:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/api/v1/userRating", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")
			ctx.Set(th.userIDKey, th.userIDValue)

			if th.paramExists == true {
				ctx.QueryParams().Set("movie_id", th.param)
			}

			th.mock()

			handler := NewProfileHandler(logger, mockService)
			handler.Register(server)

			rating := handler.GetRating()
			_ = rating(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}
