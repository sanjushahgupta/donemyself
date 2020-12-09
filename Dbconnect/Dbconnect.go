package Dbconnect

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser1"
	password = "mypass1"
	dbname   = "firstattemp"
)

func Openconnection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database

	db, _ := sql.Open("postgres", psqlconn)

	return db
}
