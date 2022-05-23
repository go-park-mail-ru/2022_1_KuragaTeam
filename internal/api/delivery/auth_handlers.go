package delivery

import (
	"context"
	"github.com/mailru/easyjson"
	"myapp/internal/constants"
	authorization "myapp/internal/microservices/authorization/proto"
	"myapp/internal/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stroiman/go-automapper"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authHandler struct {
	logger *zap.SugaredLogger

	authMicroservice authorization.AuthorizationClient
}

func NewAuthHandler(logger *zap.SugaredLogger, auth authorization.AuthorizationClient) *authHandler {
	return &authHandler{authMicroservice: auth, logger: logger}
}

func (a *authHandler) Register(router *echo.Echo) {
	router.POST(constants.SignupURL, a.SignUp())
	router.POST(constants.LoginURL, a.LogIn())
	router.DELETE(constants.LogoutURL, a.LogOut())
}

func (a *authHandler) ParseError(ctx echo.Context, requestID string, err error) error {
	if getErr, ok := status.FromError(err); ok == true {
		switch getErr.Code() {
		case codes.Internal:
			a.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		case codes.NotFound:
			a.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusNotFound),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusNotFound,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusNotFound, resp)
		case codes.InvalidArgument:
			a.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusBadRequest,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusBadRequest, resp)
		case codes.Unavailable:
			a.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

	}
	return nil
}

func (a *authHandler) LogIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := models.LogInUserDTO{}

		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			a.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		if err := ctx.Bind(&userData); err != nil {
			a.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		data := &authorization.LogInData{}
		automapper.MapLoose(userData, data)

		session, err := a.authMicroservice.LogIn(context.Background(), data)
		if err != nil {
			return a.ParseError(ctx, requestID, err)
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session.Cookie,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		a.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.UserCanBeLoggedIn,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (a *authHandler) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := models.CreateUserDTO{}

		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			a.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		if err := ctx.Bind(&userData); err != nil {
			a.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		data := &authorization.SignUpData{}
		automapper.MapLoose(userData, data)
		session, err := a.authMicroservice.SignUp(context.Background(), data)
		if err != nil {
			return a.ParseError(ctx, requestID, err)
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session.Cookie,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		a.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusCreated),
		)

		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.UserCreated,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (a *authHandler) LogOut() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			a.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			a.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		data := &authorization.Cookie{Cookie: cookie.Value}
		_, err = a.authMicroservice.LogOut(context.Background(), data)
		if err != nil {
			return a.ParseError(ctx, requestID, err)
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		ctx.SetCookie(cookie)

		a.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.UserIsLoggedOut,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}
