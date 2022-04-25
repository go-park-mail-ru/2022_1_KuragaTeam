package delivery

import (
	"context"
	"myapp/internal/constants"
	"myapp/internal/csrf"
	authorization "myapp/internal/microservices/authorization/proto"
	profile "myapp/internal/microservices/profile/proto"
	"myapp/internal/models"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stroiman/go-automapper"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type APIMicroservices struct {
	logger *zap.SugaredLogger

	authMicroservice    authorization.AuthorizationClient
	profileMicroservice profile.ProfileClient
}

func NewAPIMicroservices(logger *zap.SugaredLogger, auth authorization.AuthorizationClient,
	profile profile.ProfileClient) APIMicroservices {
	return APIMicroservices{
		logger:              logger,
		authMicroservice:    auth,
		profileMicroservice: profile,
	}
}

func (api *APIMicroservices) Register(router *echo.Echo) {
	//authorization
	router.POST(constants.SignupURL, api.SignUp())
	router.POST(constants.LoginURL, api.LogIn())
	router.DELETE(constants.LogoutURL, api.LogOut())

	//profile
	router.GET(constants.ProfileURL, api.GetUserProfile())
	router.PUT(constants.EditURL, api.EditProfile())
	router.PUT(constants.AvatarURL, api.EditAvatar())
	router.GET(constants.CsrfURL, api.GetCsrf())
	router.GET(constants.AuthURL, api.Auth())
	router.POST(constants.AddLikeUrl, api.AddLike())
	router.DELETE(constants.RemoveLikeUrl, api.RemoveLike())
	router.GET(constants.FavoritesUrl, api.GetFavorites())
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
		case codes.Unavailable:
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
		}

	}
	return nil
}

func (api *APIMicroservices) LogIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userData := models.LogInUserDTO{}

		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

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

		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

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
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
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

func (api *APIMicroservices) Auth() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		avatarName := strings.ReplaceAll(ctx.Request().Header.Get("Req"), "/api/v1/avatars/", "")

		data := &profile.UserID{ID: userID}
		userAvatar, err := api.profileMicroservice.GetAvatar(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

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

		if avatarName != userAvatar.Name {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", "wrong avatar"),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusForbidden, &models.Response{
				Status:  http.StatusForbidden,
				Message: "wrong avatar",
			})
		}

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
	}
}

func (api *APIMicroservices) GetUserProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		data := &profile.UserID{ID: userID}
		userData, err := api.profileMicroservice.GetUserProfile(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

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

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		profileData := models.ProfileUserDTO{
			Name:   userData.Name,
			Email:  userData.Email,
			Avatar: userData.Avatar,
		}

		sanitizer := bluemonday.UGCPolicy()
		profileData.Name = sanitizer.Sanitize(profileData.Name)

		return ctx.JSON(http.StatusOK, &models.ResponseUserProfile{
			Status:   http.StatusOK,
			UserData: &profileData,
		})
	}
}

func (api *APIMicroservices) EditAvatar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		file, err := ctx.FormFile("file")
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

		src, err := file.Open()
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

		buffer := make([]byte, file.Size)
		_, err = src.Read(buffer)
		src.Close()

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

		file, err = ctx.FormFile("file")
		src, err = file.Open()
		defer src.Close()
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

		fileType := http.DetectContentType(buffer)

		// Validate File Type
		if _, ex := constants.IMAGE_TYPES[fileType]; !ex {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.FileTypeIsNotSupported),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			return ctx.JSON(http.StatusBadRequest, &models.Response{
				Status:  http.StatusBadRequest,
				Message: constants.FileTypeIsNotSupported,
			})
		}

		uploadData := &profile.UploadInputFile{
			ID:          userID,
			File:        buffer,
			Size:        file.Size,
			ContentType: fileType,
		}

		fileName, err := api.profileMicroservice.UploadAvatar(context.Background(), uploadData)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		editData := &profile.EditAvatarData{
			ID:     userID,
			Avatar: fileName.Name,
		}

		_, err = api.profileMicroservice.EditAvatar(context.Background(), editData)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
	}
}

func (api *APIMicroservices) EditProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		userData := models.EditProfileDTO{}

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

		data := &profile.EditProfileData{
			ID:       userID,
			Name:     userData.Name,
			Password: userData.Password,
		}

		_, err := api.profileMicroservice.EditProfile(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
	}
}

func (api *APIMicroservices) GetCsrf() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
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

		token, err := csrf.Tokens.Create(cookie.Value, time.Now().Add(time.Hour).Unix())

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

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: token,
		})
	}
}

func (api *APIMicroservices) AddLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		movieID := models.LikeDTO{}

		if err := ctx.Bind(&movieID); err != nil {
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

		data := &profile.LikeData{
			UserID:  userID,
			MovieID: int64(movieID.ID),
		}

		_, err := api.profileMicroservice.AddLike(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsEdited,
		})
	}
}

func (api *APIMicroservices) RemoveLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		movieID := models.LikeDTO{}

		if err := ctx.Bind(&movieID); err != nil {
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

		data := &profile.LikeData{
			UserID:  userID,
			MovieID: int64(movieID.ID),
		}

		_, err := api.profileMicroservice.RemoveLike(context.Background(), data)
		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsRemoved,
		})
	}
}

func (api *APIMicroservices) GetFavorites() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			api.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			api.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.SessionRequired,
			})
		}

		if userID == -1 {
			api.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		data := &profile.UserID{ID: userID}
		userData, err := api.profileMicroservice.GetFavorites(context.Background(), data)

		if err != nil {
			return api.ParseError(ctx, requestID, err)
		}

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

		api.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		responseData := &models.FavoritesID{ID: userData.MovieId}

		return ctx.JSON(http.StatusOK, &models.ResponseFavorites{
			Status:        http.StatusOK,
			FavoritesData: responseData,
		})
	}
}
