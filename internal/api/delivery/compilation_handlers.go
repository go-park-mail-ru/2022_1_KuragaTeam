package delivery

import (
	"myapp/internal"
	"myapp/internal/constants"
	compilations "myapp/internal/microservices/compilations/proto"
	profileMicroservice "myapp/internal/microservices/profile/proto"
	"myapp/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	MCAllMoviesURL = "api/v1/movies"
	MCAllSeriesURL = "api/v1/series"
	MCByPersonURL  = "api/v1/movieCompilations/person/:person_id"
	MCByMovieURL   = "api/v1/movieCompilations/movie/:movie_id"
	MCByGenreURL   = "api/v1/movieCompilations/genre/:genre_id"
	MCTopURL       = "api/v1/movieCompilations/top"
	MCYearTopURL   = "api/v1/movieCompilations/yearTop/:year"
	MCDefaultURL   = "/api/v1/movieCompilations"
	MCFavoritesURL = "/api/v1/favorites"
	MCFindURL      = "/api/v1/find"
)

type compilationsHandler struct {
	compilationsMicroservice compilations.MovieCompilationsClient
	profileMicroservice      profileMicroservice.ProfileClient
	logger                   *zap.SugaredLogger
}

func NewCompilationsHandler(service compilations.MovieCompilationsClient, profile profileMicroservice.ProfileClient,
	logger *zap.SugaredLogger) *compilationsHandler {
	return &compilationsHandler{compilationsMicroservice: service, logger: logger,
		profileMicroservice: profile}
}

func (h *compilationsHandler) Register(router *echo.Echo) {
	router.GET(MCAllMoviesURL, h.GetAllMovies())
	router.GET(MCAllSeriesURL, h.GetAllSeries())
	router.GET(MCDefaultURL, h.GetMoviesCompilations())
	router.GET(MCByMovieURL, h.GetMCByMovieID())
	router.GET(MCByGenreURL, h.GetMCByGenre())
	router.GET(MCByPersonURL, h.GetMCByPersonID())
	router.GET(MCTopURL, h.GetTopMC())
	router.GET(MCYearTopURL, h.GetYearTopMC())
	router.GET(MCFavoritesURL, h.GetFavorites())
	router.POST(MCFindURL, h.Find())
}

func convertMC(in *compilations.MovieCompilation) *internal.MovieCompilation {
	returnMC := internal.MovieCompilation{
		Name: in.Name,
	}

	for _, movie := range in.Movies {
		returnMovie := internal.MovieInfo{
			ID:      int(movie.ID),
			Name:    movie.Name,
			Picture: movie.Picture,
		}
		for _, genre := range movie.Genre {
			returnMovie.Genre = append(returnMovie.Genre, internal.Genre{
				ID:   int(genre.ID),
				Name: genre.Name,
			})
		}
		returnMC.Movies = append(returnMC.Movies, returnMovie)
	}
	return &returnMC
}

func convertSearchCompilations(in *compilations.SearchCompilation) *models.SearchCompilation {
	returnMC := models.SearchCompilation{}

	for _, movie := range in.Movies {
		returnMovie := models.MovieInfo{
			ID:      int(movie.ID),
			Name:    movie.Name,
			Picture: movie.Picture,
		}
		for _, genre := range movie.Genre {
			returnMovie.Genre = append(returnMovie.Genre, models.Genre{
				ID:   int(genre.ID),
				Name: genre.Name,
			})
		}
		returnMC.Movies = append(returnMC.Movies, returnMovie)
	}

	for _, series := range in.Series {
		returnSeries := models.MovieInfo{
			ID:      int(series.ID),
			Name:    series.Name,
			Picture: series.Picture,
		}
		for _, genre := range series.Genre {
			returnSeries.Genre = append(returnSeries.Genre, models.Genre{
				ID:   int(genre.ID),
				Name: genre.Name,
			})
		}
		returnMC.Series = append(returnMC.Series, returnSeries)
	}

	for _, person := range in.Persons {
		returnPersons := models.PersonInfo{
			ID:       int(person.ID),
			Name:     person.Name,
			Photo:    person.Photo,
			Position: person.Position,
		}
		returnMC.Persons = append(returnMC.Persons, returnPersons)
	}
	return &returnMC
}

func (h *compilationsHandler) GetAllMovies() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		selectedMC, err := h.compilationsMicroservice.GetAllMovies(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit), Offset: int32(offset)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, convertMC(selectedMC).Movies)
	}
}

func (h *compilationsHandler) GetAllSeries() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		selectedMC, err := h.compilationsMicroservice.GetAllSeries(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit), Offset: int32(offset)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC).Movies)
	}
}

func (h *compilationsHandler) GetMoviesCompilations() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		mainMoviesCompilations, err := h.compilationsMicroservice.GetMainCompilations(context.Background(), &compilations.GetMainCompilationsOptions{})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		var returnMCs []internal.MovieCompilation

		for _, MC := range mainMoviesCompilations.MovieCompilations {
			returnMCs = append(returnMCs, *convertMC(MC))
		}

		return ctx.JSON(http.StatusOK, returnMCs)
	}
}

func (h *compilationsHandler) GetMCByMovieID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		movieID, err := strconv.Atoi(ctx.Param("movie_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.compilationsMicroservice.GetByMovie(context.Background(), &compilations.GetByIDOptions{ID: int32(movieID)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC))
	}
}

func (h *compilationsHandler) GetMCByGenre() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		genreID, err := strconv.Atoi(ctx.Param("genre_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.compilationsMicroservice.GetByGenre(context.Background(), &compilations.GetByIDOptions{ID: int32(genreID)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC))
	}
}

func (h *compilationsHandler) GetMCByPersonID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		personID, err := strconv.Atoi(ctx.Param("person_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.compilationsMicroservice.GetByPerson(context.Background(), &compilations.GetByIDOptions{ID: int32(personID)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC))
	}
}

func (h *compilationsHandler) GetTopMC() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		var limit int
		echo.QueryParamsBinder(ctx).Int("limit", &limit)
		if limit > 12 {
			limit = 12
		}
		selectedMC, err := h.compilationsMicroservice.GetTop(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC))
	}
}

func (h *compilationsHandler) GetYearTopMC() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.compilationsMicroservice.GetTopByYear(context.Background(), &compilations.GetByIDOptions{ID: int32(year)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, convertMC(selectedMC))
	}
}

func (h *compilationsHandler) GetFavorites() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			h.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			h.logger.Error(
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
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		data := &profileMicroservice.UserID{ID: userID}
		favorites, err := h.profileMicroservice.GetFavorites(context.Background(), data)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		moviesData := &compilations.GetFavoritesOptions{
			Id:     favorites.MovieId,
			Limit:  int64(limit),
			Offset: int64(offset),
		}
		selectedMC, err := h.compilationsMicroservice.GetFavorites(context.Background(), moviesData)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return ctx.JSON(http.StatusOK, convertMC(selectedMC).Movies)
	}
}

func (h *compilationsHandler) Find() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			h.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			h.logger.Error(
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
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			return ctx.JSON(http.StatusUnauthorized, &models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
		}

		movieID := models.FindDTO{}

		if err := ctx.Bind(&movieID); err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		data := &compilations.SearchText{Text: movieID.Text}

		searchedMC, err := h.compilationsMicroservice.Find(context.Background(), data)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, convertSearchCompilations(searchedMC))
	}
}
