package usecase

import (
	"myapp/internal"
	"myapp/internal/persons"
	"myapp/internal/position"
)

type service struct {
	staffStorage    persons.Storage
	positionStorage position.Storage
}

func NewService(staffStorage persons.Storage, positionStorage position.Storage) persons.Service {
	return &service{staffStorage: staffStorage, positionStorage: positionStorage}
}

func (s *service) GetByID(id int) (*internal.Person, error) {
	person, err := s.staffStorage.GetByPersonID(id)
	if err != nil {
		return nil, err
	}
	person.Position, err = s.positionStorage.GetByPersonID(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}
