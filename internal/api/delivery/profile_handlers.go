package delivery

import (
	"context"
	"github.com/mailru/easyjson"
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
	router.POST(constants.AddLikeURL, p.AddLike())
	router.DELETE(constants.RemoveLikeURL, p.RemoveLike())
	router.GET(constants.LikesURL, p.GetFavorites())
	router.GET(constants.UserRatingURL, p.GetRating())
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

			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		case codes.Unavailable:
			p.logger.Info(
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
		case codes.InvalidArgument:
			p.logger.Info(
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
		case codes.PermissionDenied:
			p.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusForbidden),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusForbidden,
				Message: getErr.Message(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusForbidden, resp)
		}

	}
	return nil
}

func (p *profileHandler) Auth() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		avatarName := strings.ReplaceAll(ctx.Request().Header.Get("Req"), "/api/v1/minio/avatars/", "")

		data := &profile.UserID{ID: userID}
		userAvatar, err := p.profileMicroservice.GetAvatar(context.Background(), data)

		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		if avatarName != userAvatar.Name {
			return constants.RespError(ctx, p.logger, requestID, "wrong avatar", http.StatusForbidden)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) Check() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		data := &profile.UserID{ID: userID}
		_, err = p.profileMicroservice.IsSubscription(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: "ok",
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) GetUserProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
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
		resp, err := easyjson.Marshal(&models.ResponseUserProfile{
			Status:   http.StatusOK,
			UserData: &profileData,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) EditAvatar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		file, err := ctx.FormFile("file")
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		src, err := file.Open()
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		buffer := make([]byte, file.Size)
		_, err = src.Read(buffer)
		err = src.Close()
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		file, err = ctx.FormFile("file")
		src, err = file.Open()
		defer src.Close()
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		fileType := http.DetectContentType(buffer)

		// Validate File Type
		if _, ex := constants.ImageTypes[fileType]; !ex {
			return constants.RespError(ctx, p.logger, requestID, constants.FileTypeIsNotSupported, http.StatusBadRequest)
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
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) EditProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		userData := models.EditProfileDTO{}

		if err = ctx.Bind(&userData); err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusBadRequest)
		}

		data := &profile.EditProfileData{
			ID:       userID,
			Name:     userData.Name,
			Password: userData.Password,
		}

		_, err = p.profileMicroservice.EditProfile(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.ProfileIsEdited,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) GetCsrf() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			return constants.RespError(ctx, p.logger, requestID, constants.NoRequestID, http.StatusInternalServerError)
		}

		cookie, err := ctx.Cookie("Session_cookie")
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		token, err := csrf.Tokens.Create(cookie.Value, time.Now().Add(time.Hour).Unix())

		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: token,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) AddLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		movieID := models.LikeDTO{}

		if err = ctx.Bind(&movieID); err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		data := &profile.LikeData{
			UserID:  userID,
			MovieID: int64(movieID.ID),
		}

		_, err = p.profileMicroservice.AddLike(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsEdited,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) RemoveLike() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		movieID := models.LikeDTO{}

		if err = ctx.Bind(&movieID); err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		data := &profile.LikeData{
			UserID:  userID,
			MovieID: int64(movieID.ID),
		}

		_, err = p.profileMicroservice.RemoveLike(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.LikeIsRemoved,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) GetFavorites() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
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

		resp, err := easyjson.Marshal(&models.ResponseFavorites{
			Status:        http.StatusOK,
			FavoritesData: responseData,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) GetPaymentsToken() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		data := &profile.UserID{ID: userID}
		token, err := p.profileMicroservice.GetPaymentsToken(context.Background(), data)
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: token.Token,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) Payment() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		token := models.TokenDTO{}

		if err := ctx.Bind(&token); err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
		}

		data := &profile.CheckTokenData{
			Token: token.Token,
			Id:    userID,
		}
		_, err = p.profileMicroservice.CheckPaymentsToken(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}

		_, err = p.profileMicroservice.CreatePayment(context.Background(), data)
		if err != nil {
			return p.ParseError(ctx, requestID, err)
		}
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.PaymentIsCreated,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) Subscribe() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			return constants.RespError(ctx, p.logger, requestID, constants.NoRequestID, http.StatusInternalServerError)
		}

		headerContentType := ctx.Request().Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			return constants.RespError(ctx, p.logger, requestID, constants.UnsupportedMediaType, http.StatusUnsupportedMediaType)
		}

		if err := ctx.Request().ParseForm(); err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusUnsupportedMediaType)
		}

		payToken := ctx.Request().PostForm["label"][0]
		amount := ctx.Request().PostForm["withdraw_amount"][0]

		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, err.Error(), http.StatusInternalServerError)
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
		resp, err := easyjson.Marshal(&models.Response{
			Status:  http.StatusOK,
			Message: constants.PaymentIsCreated,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (p *profileHandler) GetRating() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, requestID, err := constants.DefaultUserChecks(ctx, p.logger)
		if err != nil {
			return err
		}

		movieIDStr := ctx.QueryParam("movie_id")
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			return constants.RespError(ctx, p.logger, requestID, constants.NoMovieID, http.StatusBadRequest)
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

			resp, err := easyjson.Marshal(&models.ResponseMovieRating{
				Status: http.StatusOK,
				Rating: int(userRating.Rating),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusOK, resp)
		}

		p.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		resp, err := easyjson.Marshal(&models.ResponseMovieRating{
			Status: http.StatusNotFound,
		})
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusNotFound, resp)
	}
}
