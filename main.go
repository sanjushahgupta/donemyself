package main

import (
	"database/sql"
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Contact struct {
	id   int    `json:"id"`
	Name string `json:"name"`

	Locality string `json:"Locality"`
}

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
	db.Ping()
	return db
}
func Create(w http.ResponseWriter, r *http.Request) {
	db := Openconnection()
	var data Contact
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("querytest")
	defer db.Close()
	sqlStatement := `
		INSERT INTO "newtable"("Name","Locality")
		VALUES ($1,$2)`
	_, err = db.Exec(sqlStatement, data.Name, data.Locality)

	if err != nil {
		panic(err)
	}

}

func List(w http.ResponseWriter, r *http.Request) {

	db := Openconnection()
	var newtable Contact
	var Contactsarr []Contact
	defer db.Close()

	rows, err := db.Query(`SELECT "Name", "Locality" FROM "newtable"`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&newtable.Name, &newtable.Locality)
		if err != nil {
			fmt.Println(err)
		} else {
			Contactsarr = append(Contactsarr, newtable)
		}

	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Contactsarr)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/addcontact", Create).Methods("POST")
	r.HandleFunc("/readcontact", List).Methods("GET")
	http.ListenAndServe(":8099", r)

}
