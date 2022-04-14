package delivery

import (
	"encoding/json"
	"errors"
	"myapp/internal"
	"myapp/mock"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestPersonsDelivery_GetPerson(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person := internal.Person{
		ID:          1,
		Name:        "Person1",
		Photo:       "personPhoto.webp",
		AdditPhoto1: "personAddit1Photo.webp",
		AdditPhoto2: "personAddit2Photo.webp",
		Description: "Description of person",
		Position: []string{
			"Актер",
			"Режиссер",
			"Сценарист",
		},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockPersonsService
		expected      internal.Person
		expectedError bool
	}{
		{
			name:        "Get person by ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockPersonsService{
				GetByIDFunc: func(id int) (*internal.Person, error) {
					return &person, nil
				},
			},
			expected: internal.Person{
				ID:          person.ID,
				Name:        person.Name,
				Photo:       person.Photo,
				AdditPhoto1: person.AdditPhoto1,
				AdditPhoto2: person.AdditPhoto2,
				Description: person.Description,
				Position:    person.Position,
			},
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockPersonsService{
				GetByIDFunc: func(id int) (*internal.Person, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:          "No param",
			paramExists:   false,
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/person/0", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			if test.paramExists {
				ctx.SetParamNames("person_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock, logger)
			r.Register(server)
			personFromDelivery := r.GetPerson()

			_ = personFromDelivery(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result internal.Person
			_ = json.Unmarshal(body.Bytes(), &result)
			if test.expectedError {
				assert.Equal(t, "500 Internal Server Error", status)
			} else {
				assert.Equal(t, test.expected, result)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}
