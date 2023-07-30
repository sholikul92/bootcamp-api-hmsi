package connectdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnPostgress(dbHost, dbPort, dbUser, dbPassword, dbName, dbDriver string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open(dbDriver, psqlInfo)

	if err != nil {
		return nil, err
	}
	// defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
