package delivery

import (
	"bytes"
	"mime/multipart"
	"myapp/internal/constants"
	"myapp/internal/csrf"
	"myapp/internal/microservices/profile/proto"
	"myapp/internal/microservices/profile/usecase"
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
				}
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetUserProfile(gomock.Any(), inputData).Return(userData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"user\":{\"username\":\"Olga\",\"email\":\"olga@mail.ru\",\"avatar\":\"avatar\"}}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(-1),
			fieldName:      "file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, Error file",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"http: no such file\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			fieldName:      "wrong_file",
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, Error content type",
			mock:           func() {},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"File type is not supported\"}\n",
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
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"username": "Olga", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":200,\"message\":\"Profile is edited\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
			expectedJSON:   "{\"status\":200,\"message\":\"" + token + "\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"http: named cookie not present\"}\n",
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
			expectedJSON:   "{\"status\":200,\"message\":\"ok\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"No RequestID in context\"}\n",
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
			expectedJSON:   "{\"status\":403,\"message\":\"wrong avatar\"}\n",
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
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":200,\"message\":\"" + constants.LikeIsEdited + "\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(1),
			data:           `{"id": "3"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":200,\"message\":\"" + constants.LikeIsRemoved + "\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
				favorite := &proto.Favorites{MovieId: likes}
				inputData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().GetFavorites(gomock.Any(), inputData).Return(favorite, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"favorites\":{\"id\":[1,2,3]}}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, ctx hasn't USER_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"Session required\"}\n",
			userIDKey:      "WRONG_USER_ID",
			userIDValue:    int64(-1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong REQUEST_ID",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
			userIDKey:      "USER_ID",
			userIDValue:    int64(1),
			requestID:      "REQUEST_ID",
		},
		{
			name:           "Handler returned status 401, User is unauthorized",
			mock:           func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedJSON:   "{\"status\":401,\"message\":\"User is unauthorized\"}\n",
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
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}\n",
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
