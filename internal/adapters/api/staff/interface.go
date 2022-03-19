package staff

import "myapp/internal/domain"

type Service interface {
	GetByID(id int) (*domain.Person, error)
}
