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

func (r *customerRepository) Insert(c *models.RequestInsertCustomer) error {
	err := r.Repo.Create(c)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) FindById(id uint) (*models.Customers, error) {
	customer, err := r.Repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) UpdateById(id uint, c *models.RequestInsertCustomer) error {
	exist, err := r.Repo.GetById(id)
	if err != nil {
		return err
	}

	if exist == nil {
		return err
	}

	return r.Repo.Update(id, c)
}
