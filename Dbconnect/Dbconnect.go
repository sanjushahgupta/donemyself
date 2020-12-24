package Dbconnect

import (
	"firstattemp/Model"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser1"
	password = "mypass1"
	dbname   = "firstattemp"
)

func Openconnection() *gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database

	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Model.User{}, &Model.Jobdetails{}, &Model.Token{})
	return db
}
