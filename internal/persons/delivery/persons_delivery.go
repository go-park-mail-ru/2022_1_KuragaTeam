package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"myapp/internal/persons"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	personURL = "api/v1/person/:person_id"
)

type handler struct {
	staffService persons.Service
	logger       *zap.SugaredLogger
}

func NewHandler(service persons.Service, logger *zap.SugaredLogger) *handler {
	return &handler{staffService: service, logger: logger}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(personURL, h.GetPerson())
}

func (h *handler) GetPerson() echo.HandlerFunc {
	return func(context echo.Context) error {
		requestID := context.Get("REQUEST_ID").(string)
		personID, err := strconv.Atoi(context.Param("person_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedPerson, err := h.staffService.GetByID(personID)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return context.JSON(http.StatusOK, &selectedPerson)
	}
}
