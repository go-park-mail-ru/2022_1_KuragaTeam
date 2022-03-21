package staff

import (
	"myapp/internal"
	"myapp/internal/adapters/api/staff"
	"myapp/internal/domain/position"
)

type service struct {
	staffStorage    Storage
	positionStorage position.Storage
}

func NewService(staffStorage Storage, positionStorage position.Storage) staff.Service {
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
