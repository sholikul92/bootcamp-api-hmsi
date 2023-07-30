package query

import (
	"database/sql"
	"fmt"
)

type Customers struct {
	Id    uint
	Name  string
	Phone string
	Email string
	Age   uint
}

type DB struct {
	Conn *sql.DB
}

// Create customers
func (db *DB) Create(c *Customers) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		fmt.Println("error :", err)
		return err
	}
	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)
	if err != nil {
		return errExec
	}
	return nil
}

// Read customers
func (db *DB) Read() (*[]Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT * FROM customers`)
	if errExec != nil {
		return nil, errExec
	}
	var result []Customers
	for rows.Next() {
		var each Customers

		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)
		if errScan != nil {
			return nil, errScan
		}
		result = append(result, each)
	}

	return &result, nil
}
