package customers

import (
	"bootcamp-api-hmsi/models"
)

type (
	CustomersRepository interface {
		GetAll() (*[]models.Customers, error)
	}

	CustomersUseCase interface {
		FindAll() (*[]models.Customers, error)
	}
)
