package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/utils/images"
	mock2 "myapp/mock"
	"testing"
)

func TestPersonsUsecase_GetByID(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person := internal.Person{
		ID:          1,
		Name:        "Актер 1",
		Photo:       "photo.webp",
		AdditPhoto1: "additPhoto1.webp",
		AdditPhoto2: "additPhoto1.webp",
		Description: "Description",
		Position:    []string{"Актер", "Сценарист"},
	}

	personFromStorage := internal.Person{
		ID:          person.ID,
		Name:        person.Name,
		Photo:       person.Photo,
		AdditPhoto1: person.AdditPhoto1,
		AdditPhoto2: person.AdditPhoto2,
		Description: person.Description,
		Position:    person.Position,
	}

	person.Photo, _ = images.GenerateFileURL(person.Photo, "persons")
	person.AdditPhoto1, _ = images.GenerateFileURL(person.AdditPhoto1, "persons")
	person.AdditPhoto2, _ = images.GenerateFileURL(person.AdditPhoto2, "persons")

	tests := []struct {
		name                string
		personStorageMock   *mock2.MockPersonsStorage
		positionStorageMock *mock2.MockPositionStorage
		expected            internal.Person
		expectedError       bool
	}{
		{
			name: "Get one person",
			personStorageMock: &mock2.MockPersonsStorage{
				GetByPersonIDFunc: func(id int) (*internal.Person, error) {
					return &personFromStorage, nil
				},
			},
			positionStorageMock: &mock2.MockPositionStorage{
				GetByPersonIDFunc: func(id int) ([]string, error) {
					return personFromStorage.Position, nil
				},
			},
			expected:      person,
			expectedError: false,
		},
		{
			name: "Persons storage error",
			personStorageMock: &mock2.MockPersonsStorage{
				GetByPersonIDFunc: func(id int) (*internal.Person, error) {
					return nil, errors.New(testError)
				},
			},
			positionStorageMock: &mock2.MockPositionStorage{
				GetByPersonIDFunc: func(id int) ([]string, error) {
					return personFromStorage.Position, nil
				},
			},
			expectedError: true,
		},

		{
			name: "Positions storage error",
			personStorageMock: &mock2.MockPersonsStorage{
				GetByPersonIDFunc: func(id int) (*internal.Person, error) {
					return &personFromStorage, nil
				},
			},
			positionStorageMock: &mock2.MockPositionStorage{
				GetByPersonIDFunc: func(id int) ([]string, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.personStorageMock, test.positionStorageMock)
			mainMovie, err := r.GetByID(1)

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, *mainMovie)
			}
		})
	}
}
