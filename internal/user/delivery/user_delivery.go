package delivery

import (
	"myapp/internal/csrf"
	"myapp/internal/user"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

const (
	signupURL  = "/api/v1/signup"
	loginURL   = "/api/v1/login"
	logoutURL  = "/api/v1/logout"
	profileURL = "/api/v1/profile"
	editURL    = "/api/v1/edit"
	avatarURL  = "/api/v1/avatar"
	csrfURL    = "/api/v1/csrf"
	authURL    = "/api/v1/auth"
)

const (
	UserNotFound           = "User not found"
	UserCanBeLoggedIn      = "User can be logged in"
	UserCreated            = "User created"
	SessionRequired        = "Session required"
	UserIsUnauthorized     = "User is unauthorized"
	UserIsLoggedOut        = "User is logged out"
	FileTypeIsNotSupported = "File type is not supported"
	ProfileIsEdited        = "Profile is edited"
)

var (
	IMAGE_TYPES = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}
)

type handler struct {
	userService user.Service
	logger      *zap.SugaredLogger
}

func NewHandler(service user.Service, logger *zap.SugaredLogger) *handler {
	return &handler{
		userService: service,
		logger:      logger,
	}
}

func (h *handler) Register(router *echo.Echo) {
	router.POST(signupURL, h.SignUp())
	router.POST(loginURL, h.LogIn())
	router.DELETE(logoutURL, h.LogOut())
	router.GET(profileURL, h.GetUserProfile())
	router.PUT(editURL, h.EditProfile())
	router.PUT(avatarURL, h.EditAvatar())
	router.GET(csrfURL, h.GetCsrf())
	router.GET(authURL, h.Auth())
}

func (h *handler) Auth() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		userID, ok := ctx.Get("USER_ID").(int64)

		if !ok {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: SessionRequired,
			})
		}

		if userID == -1 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &user.Response{
				Status:  http.StatusUnauthorized,
				Message: UserIsUnauthorized,
			})
		}

		avatarName := strings.ReplaceAll(ctx.Request().Header.Get("Req"), "/avatars/", "")

		userAvatar, err := h.userService.GetAvatar(userID)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if avatarName != userAvatar {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", "wrong avatar"),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusForbidden, &user.Response{
				Status:  http.StatusForbidden,
				Message: "wrong avatar",
			})
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
	}
}

func (h *handler) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := user.CreateUserDTO{}
		requestID := ctx.Get("REQUEST_ID").(string)

		if err := ctx.Bind(&userData); err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		session, msg, err := h.userService.SignUp(&userData)

		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if len(session) == 0 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", msg),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			return ctx.JSON(http.StatusBadRequest, &user.Response{
				Status:  http.StatusBadRequest,
				Message: msg,
			})
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusCreated),
		)

		return ctx.JSON(http.StatusCreated, &user.Response{
			Status:  http.StatusCreated,
			Message: UserCreated,
		})
	}
}

func (h *handler) LogIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := user.LogInUserDTO{}
		requestID := ctx.Get("REQUEST_ID").(string)

		if err := ctx.Bind(&userData); err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		session, err := h.userService.LogIn(&userData)

		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if len(session) == 0 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", UserNotFound),
				zap.Int("ANSWER STATUS", http.StatusNotFound),
			)
			return ctx.JSON(http.StatusNotFound, &user.Response{
				Status:  http.StatusNotFound,
				Message: UserNotFound,
			})
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: UserCanBeLoggedIn,
		})
	}
}

func (h *handler) GetUserProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		userID, ok := ctx.Get("USER_ID").(int64)

		if !ok {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: SessionRequired,
			})
		}

		if userID == -1 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &user.Response{
				Status:  http.StatusUnauthorized,
				Message: UserIsUnauthorized,
			})
		}

		userData, err := h.userService.GetUserProfile(userID)

		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		sanitizer := bluemonday.UGCPolicy()
		userData.Avatar = sanitizer.Sanitize(userData.Avatar)
		userData.Name = sanitizer.Sanitize(userData.Name)
		userData.Email = sanitizer.Sanitize(userData.Email)

		return ctx.JSON(http.StatusOK, &user.ResponseUserProfile{
			Status:   http.StatusOK,
			UserData: userData,
		})
	}
}

func (h *handler) LogOut() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		err = h.userService.LogOut(cookie.Value)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		ctx.SetCookie(cookie)

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: UserIsLoggedOut,
		})
	}
}

func (h *handler) EditAvatar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: SessionRequired,
			})
		}

		if userID == -1 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &user.Response{
				Status:  http.StatusUnauthorized,
				Message: UserIsUnauthorized,
			})
		}

		userData := user.EditAvatarDTO{
			ID: userID,
		}

		fileName := ""

		file, err := ctx.FormFile("file")
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		src, err := file.Open()
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		buffer := make([]byte, file.Size)
		_, err = src.Read(buffer)
		src.Close()
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		file, err = ctx.FormFile("file")
		src, err = file.Open()
		defer src.Close()
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		fileType := http.DetectContentType(buffer)

		// Validate File Type
		if _, ex := IMAGE_TYPES[fileType]; !ex {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", FileTypeIsNotSupported),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			return ctx.JSON(http.StatusBadRequest, &user.Response{
				Status:  http.StatusBadRequest,
				Message: FileTypeIsNotSupported,
			})
		}

		fileName, err = h.userService.UploadAvatar(src, file.Size, fileType, userID)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		userData.Avatar = fileName

		err = h.userService.EditAvatar(&userData)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: ProfileIsEdited,
		})
	}
}

func (h *handler) EditProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: SessionRequired,
			})
		}

		if userID == -1 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &user.Response{
				Status:  http.StatusUnauthorized,
				Message: UserIsUnauthorized,
			})
		}

		userData := user.EditProfileDTO{
			ID: userID,
		}

		if err := ctx.Bind(&userData); err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		err := h.userService.EditProfile(&userData)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: ProfileIsEdited,
		})
	}
}

func (h *handler) GetCsrf() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		token, err := csrf.Tokens.Create(cookie.Value, time.Now().Add(time.Hour).Unix())

		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &user.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &user.Response{
			Status:  http.StatusOK,
			Message: token,
		})
	}
}
