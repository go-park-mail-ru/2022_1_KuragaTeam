package delivery

import (
	"myapp/internal"
	"myapp/internal/adapters/api"
	"myapp/internal/user"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	signupURL   = "/api/v1/signup"
	loginURL    = "/api/v1/login"
	logoutURL   = "/api/v1/logout"
	mainPageURL = "/api/v1"
	profileURL  = "/api/v1/profile"
	editURL     = "/api/v1/edit"
)

type handler struct {
	userService user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.POST(signupURL, h.SignUp())
	router.POST(loginURL, h.LogIn())
	router.DELETE(logoutURL, h.LogOut())
	router.GET(mainPageURL, h.GetUserMainPage())
	router.GET(profileURL, h.GetUserProfile())
	router.PUT(editURL, h.EditProfile())
}

func (h *handler) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := internal.CreateUserDTO{}

		if err := ctx.Bind(&userData); err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		session, msg, err := h.userService.SignUp(&userData)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if len(session) == 0 {
			return ctx.JSON(http.StatusBadRequest, &internal.Response{
				Status:  http.StatusBadRequest,
				Message: msg,
			})
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session,
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		return ctx.JSON(http.StatusCreated, &internal.Response{
			Status:  http.StatusCreated,
			Message: "OK: User created",
		})
	}
}

func (h *handler) LogIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := internal.LogInUserDTO{}

		if err := ctx.Bind(&userData); err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		session, err := h.userService.LogIn(&userData)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if len(session) == 0 {
			return ctx.JSON(http.StatusNotFound, &internal.Response{
				Status:  http.StatusNotFound,
				Message: "ERROR: User not found",
			})
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session,
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		return ctx.JSON(http.StatusOK, &internal.Response{
			Status:  http.StatusOK,
			Message: "OK: User can be logged in",
		})
	}
}

func (h *handler) GetUserMainPage() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: "ERROR: Session required",
			})
		}

		if userID == -1 {
			return ctx.JSON(http.StatusUnauthorized, &internal.Response{
				Status:  http.StatusUnauthorized,
				Message: "ERROR: User is unauthorized",
			})
		}

		userData, err := h.userService.GetUserMainPage(userID)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, &internal.ResponseUserMainPage{
			Status:   http.StatusOK,
			UserData: userData,
		})
	}
}

func (h *handler) GetUserProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: "ERROR: Session required",
			})
		}

		if userID == -1 {
			return ctx.JSON(http.StatusUnauthorized, &internal.Response{
				Status:  http.StatusUnauthorized,
				Message: "ERROR: User is unauthorized",
			})
		}

		userData, err := h.userService.GetUserProfile(userID)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, &internal.ResponseUserProfile{
			Status:   http.StatusOK,
			UserData: userData,
		})
	}
}

func (h *handler) LogOut() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		err = h.userService.LogOut(cookie.Value)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusOK, &internal.Response{
			Status:  http.StatusOK,
			Message: "OK: User is logged out",
		})
	}
}

func (h *handler) EditProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: "ERROR: Session required",
			})
		}

		if userID == -1 {
			return ctx.JSON(http.StatusUnauthorized, &internal.Response{
				Status:  http.StatusUnauthorized,
				Message: "ERROR: User is unauthorized",
			})
		}

		userData := internal.EditProfileDTO{
			ID: userID,
		}

		if err := ctx.Bind(&userData); err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if userData.Password != userData.RepeatPassword {
			return ctx.JSON(http.StatusBadRequest, &internal.Response{
				Status:  http.StatusBadRequest,
				Message: "ERROR: Passwords are different",
			})
		}

		err := h.userService.EditProfile(&userData)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, &internal.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, &internal.Response{
			Status:  http.StatusOK,
			Message: "OK: Profile is edited",
		})
	}
}
