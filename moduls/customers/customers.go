package customers

import (
	"bootcamp-api-hmsi/models"
)

type (
	CustomersRepository interface {
		Create(c *models.RequestInsertCustomer) error
		GetAll() (*[]models.Customers, error)
		GetById(id uint) (*models.Customers, error)
		Update(id uint, c *models.RequestInsertCustomer) error
	}

	CustomersUseCase interface {
		Insert(c *models.RequestInsertCustomer) error
		FindAll() (*[]models.Customers, error)
		FindById(id uint) (*models.Customers, error)
		UpdateById(id uint, c *models.RequestInsertCustomer) error
	}
)
