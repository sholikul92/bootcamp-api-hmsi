package customerRepository

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/moduls/customers"
	"database/sql"
)

type DB struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) customers.CustomersRepository {
	return &DB{Conn}
}

// Get all users
func (db *DB) GetAll() (*[]models.Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT * FROM customers`)
	if errExec != nil {
		return nil, errExec
	}

	// deklarasil variabel result untuk menampung hasil rows
	var result []models.Customers

	for rows.Next() {
		var each models.Customers

		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)

		if errScan != nil {
			return nil, errScan
		}
		result = append(result, each)
	}

	return &result, nil
}

// Create user
func (db *DB) Create(c *models.RequestInsertCustomer) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1, $2, $3, $4)`)

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)

	if errExec != nil {
		return errExec
	}

	return nil
}

// Get customer by id
func (db *DB) GetById(id uint) (*models.Customers, error) {
	query := "SELECT id, name, phone, email, age FROM customers WHERE id = $1"
	row := db.Conn.QueryRow(query, id)

	var customer models.Customers
	err := row.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Email, &customer.Age)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// update Customer
func (db *DB) Update(id uint, c *models.RequestInsertCustomer) error {
	query := "UPDATE customers SET name=$1, phone=$2, email=$3, age=$4 WHERE id=$5"

	_, err := db.Conn.Exec(query, c.Name, c.Phone, c.Email, c.Age, id)
	if err != nil {
		return err
	}

	return nil
}
