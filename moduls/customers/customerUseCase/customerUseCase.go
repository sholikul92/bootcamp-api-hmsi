package customerUseCase

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/moduls/customers"
)

type customerRepository struct {
	Repo customers.CustomersRepository
}

func NewCustomerUseCase(Repo customers.CustomersRepository) customers.CustomersUseCase {
	return &customerRepository{Repo}
}

func (r *customerRepository) FindAll() (*[]models.Customers, error) {
	result, err := r.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}
