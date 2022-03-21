package delivery

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api"
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
}

func NewHandler(service persons.Service) api.Handler {
	return &handler{staffService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(personURL, h.GetPerson())
}

func (h *handler) GetPerson() echo.HandlerFunc {
	return func(context echo.Context) error {
		personID, err := strconv.Atoi(context.Param("person_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedPerson, err := h.staffService.GetByID(personID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedPerson)
	}
}
