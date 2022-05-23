package usecase

import (
	"myapp/internal"
	"myapp/internal/microservices/profile/utils/images"
	"myapp/internal/persons"
	"myapp/internal/position"
)

type service struct {
	staffStorage    persons.Storage
	positionStorage position.Storage
}

func NewService(staffStorage persons.Storage, positionStorage position.Storage) *service {
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
	person.Photo, err = images.GenerateFileURL(person.Photo, "persons")
	if err != nil {
		return nil, err
	}
	person.AdditPhoto1, err = images.GenerateFileURL(person.AdditPhoto1, "persons")
	if err != nil {
		return nil, err
	}
	person.AdditPhoto2, err = images.GenerateFileURL(person.AdditPhoto2, "persons")
	if err != nil {
		return nil, err
	}
	return person, nil
}
