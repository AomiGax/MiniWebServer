package utils

// This package is used to put some utils, like database connection

import (
	"database/sql"
	"fmt"
)

// Database's some parameter
const (
	host     = "localhost"
	port     = 5432
	user     = "linxinyu"
	password = "linxinyu"
	dbname   = "Demo"
)

// Connect database and put the connection into the global.DB
func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
