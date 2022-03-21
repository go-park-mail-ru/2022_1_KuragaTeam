package persons

import (
	"myapp/internal"
)

type Service interface {
	GetByID(id int) (*internal.Person, error)
}
