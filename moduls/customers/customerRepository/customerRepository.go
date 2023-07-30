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
