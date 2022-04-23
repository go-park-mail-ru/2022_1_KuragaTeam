package delivery

import (
	"context"
	authorization "myapp/internal/microservices/authorization/proto"
	"myapp/internal/models"
	"myapp/internal/utils/constants"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stroiman/go-automapper"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type APIMicroservices struct {
	logger *zap.SugaredLogger

	authMicroservice authorization.AuthorizationClient
}

func NewAPIMicroservices(logger *zap.SugaredLogger, auth authorization.AuthorizationClient) APIMicroservices {
	return APIMicroservices{
		logger:           logger,
		authMicroservice: auth,
	}
}

func (api *APIMicroservices) Register(router *echo.Echo) {
	//authorization
	router.POST(constants.SignupURL, api.SignUp())
	router.POST(constants.LoginURL, api.LogIn())
	router.DELETE(constants.LogoutURL, api.LogOut())
}

func (api *APIMicroservices) ParseError(ctx echo.Context, requestID string, err error) error {
	if getErr, ok := status.FromError(err); ok == true {
		switch getErr.Code() {
		case codes.Internal:
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
		case codes.NotFound:
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusNotFound),
			)
			return ctx.JSON(http.StatusNotFound, &models.Response{
				Status:  http.StatusNotFound,
				Message: getErr.Message(),
			})
		case codes.InvalidArgument:
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			return ctx.JSON(http.StatusBadRequest, &models.Response{
				Status:  http.StatusBadRequest,
				Message: getErr.Message(),
			})

		}
	}
	return nil
}

func (api *APIMicroservices) LogIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := models.LogInUserDTO{}
		requestID := ctx.Get("REQUEST_ID").(string)

		if err := ctx.Bind(&userData); err != nil {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &authorization.LogInData{}
		automapper.MapLoose(userData, data)
		session, err := api.authMicroservice.LogIn(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session.Cookie,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.UserCanBeLoggedIn,
		})
	}
}

func (api *APIMicroservices) SignUp() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := models.CreateUserDTO{}
		requestID := ctx.Get("REQUEST_ID").(string)

		if err := ctx.Bind(&userData); err != nil {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &authorization.SignUpData{}
		automapper.MapLoose(userData, data)
		session, err := api.authMicroservice.SignUp(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		cookie := http.Cookie{
			Name:     "Session_cookie",
			Value:    session.Cookie,
			HttpOnly: true,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			SameSite: 0,
		}

		ctx.SetCookie(&cookie)

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusCreated),
		)

		return ctx.JSON(http.StatusCreated, &models.Response{
			Status:  http.StatusCreated,
			Message: constants.UserCreated,
		})
	}
}

func (api *APIMicroservices) LogOut() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.NoContent(http.StatusInternalServerError)
		}

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &authorization.Cookie{Cookie: cookie.Value}
		_, err = api.authMicroservice.LogOut(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		ctx.SetCookie(cookie)

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.UserIsLoggedOut,
		})
	}
}
