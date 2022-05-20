package delivery

import (
	"context"
	"myapp/internal/constants"
	"myapp/internal/csrf"
	profile "myapp/internal/microservices/profile/proto"
	"myapp/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type profileHandler struct {
	logger *zap.SugaredLogger

	profileMicroservice profile.ProfileClient
}

func NewProfileHandler(logger *zap.SugaredLogger, profile profile.ProfileClient) *profileHandler {
	return &profileHandler{profileMicroservice: profile, logger: logger}
}

func (p *profileHandler) Register(router *echo.Echo) {
	router.GET(constants.ProfileURL, p.GetUserProfile())
	router.PUT(constants.EditURL, p.EditProfile())
	router.PUT(constants.AvatarURL, p.EditAvatar())
	router.GET(constants.CsrfURL, p.GetCsrf())
	router.GET(constants.AuthURL, p.Auth())
	router.GET(constants.CheckURL, p.Check())
	router.POST(constants.AddLikeUrl, p.AddLike())
	router.DELETE(constants.RemoveLikeUrl, p.RemoveLike())
	router.GET(constants.LikesUrl, p.GetFavorites())
	router.GET(constants.UserRatingUrl, p.GetRating())
	router.GET(constants.PaymentsTokenURL, p.GetPaymentsToken())
	router.POST(constants.PaymentURL, p.Payment())
	router.POST(constants.SubscribeURL, p.Subscribe())
}

func (p *profileHandler) ParseError(ctx echo.Context, requestID string, err error) error {
	if getErr, ok := status.FromError(err); ok == true {
		switch getErr.Code() {
		case codes.Internal:
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
		case codes.Unavailable:
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
		case codes.InvalidArgument:
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			return ctx.JSON(http.StatusBadRequest, &models.Response{
				Status:  http.StatusBadRequest,
				Message: getErr.Message(),
			})
		case codes.PermissionDenied:
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusForbidden),
			)
			return ctx.JSON(http.StatusForbidden, &models.Response{
				Status:  http.StatusForbidden,
				Message: getErr.Message(),
			})
		}

	}
	return nil
}

func (p *profileHandler) Auth() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		avatarName := strings.ReplaceAll(ctx.Request().Header.Get("Req"), "/api/v1/minio/avatars/", "")

		data := &profile.UserID{ID: userID}
		userAvatar, err := p.profileMicroservice.GetAvatar(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		if err != nil {
			p.logger.Error(
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
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", "wrong avatar"),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusForbidden, &models.Response{
				Status:  http.StatusForbidden,
				Message: "wrong avatar",
			})
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
	}
}

func (p *profileHandler) Check() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
		_, err := p.profileMicroservice.IsSubscription(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
	}
}

func (p *profileHandler) GetUserProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
		userData, err := p.profileMicroservice.GetUserProfile(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		profileData := models.ProfileUserDTO{
			Name:   userData.Name,
			Email:  userData.Email,
			Avatar: userData.Avatar,
			Date:   userData.Date,
		}

		sanitizer := bluemonday.UGCPolicy()
		profileData.Name = sanitizer.Sanitize(profileData.Name)

		return ctx.JSON(http.StatusOK, &models.ResponseUserProfile{
			Status:   http.StatusOK,
			UserData: &profileData,
		})
	}
}

func (p *profileHandler) EditAvatar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
			p.logger.Error(
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
			p.logger.Error(
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
			p.logger.Error(
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
			p.logger.Error(
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
			p.logger.Info(
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

		fileName, err := p.profileMicroservice.UploadAvatar(context.Background(), uploadData)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		editData := &profile.EditAvatarData{
			ID:     userID,
			Avatar: fileName.Name,
		}

		_, err = p.profileMicroservice.EditAvatar(context.Background(), editData)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
	}
}

func (p *profileHandler) EditProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
			p.logger.Error(
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

		_, err := p.profileMicroservice.EditProfile(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
	}
}

func (p *profileHandler) GetCsrf() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			p.logger.Error(
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
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: token,
		})
	}
}

func (p *profileHandler) AddLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
			p.logger.Error(
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

		_, err := p.profileMicroservice.AddLike(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsEdited,
		})
	}
}

func (p *profileHandler) RemoveLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
			p.logger.Error(
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

		_, err := p.profileMicroservice.RemoveLike(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsRemoved,
		})
	}
}

func (p *profileHandler) GetFavorites() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
		userData, err := p.profileMicroservice.GetFavorites(context.Background(), data)

		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		responseData := &models.FavoritesID{ID: userData.Id}

		return ctx.JSON(http.StatusOK, &models.ResponseFavorites{
			Status:        http.StatusOK,
			FavoritesData: responseData,
		})
	}
}

func (p *profileHandler) GetPaymentsToken() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
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
		token, err := p.profileMicroservice.GetPaymentsToken(context.Background(), data)

		if err != nil {
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)

			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: token.Token,
		})
	}
}

func (p *profileHandler) Payment() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		token := models.TokenDTO{}

		if err := ctx.Bind(&token); err != nil {
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &profile.CheckTokenData{
			Token: token.Token,
			Id:    userID,
		}
		_, err := p.profileMicroservice.CheckPaymentsToken(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		_, err = p.profileMicroservice.CreatePayment(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.PaymentIsCreated,
		})
	}
}

func (p *profileHandler) Subscribe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		headerContentType := ctx.Request().Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UnsupportedMediaType),
				zap.Int("ANSWER STATUS", http.StatusUnsupportedMediaType),
			)
			return ctx.JSON(http.StatusUnsupportedMediaType, &models.Response{
				Status:  http.StatusUnsupportedMediaType,
				Message: constants.UnsupportedMediaType,
			})
		}

		if err := ctx.Request().ParseForm(); err != nil {
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusUnsupportedMediaType),
			)
			return ctx.JSON(http.StatusUnsupportedMediaType, &models.Response{
				Status:  http.StatusUnsupportedMediaType,
				Message: err.Error(),
			})
		}

		payToken := ctx.Request().PostForm["label"][0]
		amount := ctx.Request().PostForm["withdraw_amount"][0]

		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			p.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &profile.Token{
			Token: payToken,
		}
		_, err = p.profileMicroservice.CheckToken(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		subscribeData := &profile.SubscribeData{
			Token:  payToken,
			Amount: float32(amountFloat),
		}
		_, err = p.profileMicroservice.CreateSubscribe(context.Background(), subscribeData)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		return ctx.JSON(http.StatusOK, &models.Response{
			Status:  http.StatusOK,
			Message: constants.PaymentIsCreated,
		})
	}
}

func (p *profileHandler) GetRating() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			p.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			p.logger.Error(
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
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}
		movieIDStr := ctx.QueryParam("movie_id")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			p.logger.Error(
				zap.String("ERROR", constants.NoMovieId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoMovieId,
			})
		}
		data := &profile.MovieRating{UserID: userID, MovieID: int64(movieID)}

		userRating, err := p.profileMicroservice.GetMovieRating(context.Background(), data)

		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		if userRating != nil {
			p.logger.Info(
				zap.String("ID", requestID),
				zap.Int("ANSWER STATUS", http.StatusOK),
			)

			return ctx.JSON(http.StatusOK, &models.ResponseMovieRating{
				Status: http.StatusOK,
				Rating: int(userRating.Rating),
			})
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		return ctx.JSON(http.StatusNotFound, &models.ResponseMovieRating{
			Status: http.StatusNotFound,
		})
	}
}
