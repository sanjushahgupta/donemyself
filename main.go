package main

import (
	"database/sql"
	"encoding/json"
	"firstattemp/Authentication"
	"io/ioutil"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Contact struct {
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
	db9 := Openconnection()
	var data Contact
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("querytest")
	defer db9.Close()
	sqlStatement := `
		INSERT INTO "newtable"("Name","Locality")
		VALUES ($1,$2)`
	_, err = db9.Exec(sqlStatement, data.Name, data.Locality)

	if err != nil {
		panic(err)
	}

}

func List(w http.ResponseWriter, r *http.Request) {

	db2 := Openconnection()
	var newtable Contact
	var Contactsarr []Contact
	defer db2.Close()

	rows, err := db2.Query(`SELECT "Name", "Locality" FROM "newtable"`)
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

func Listbyname(w http.ResponseWriter, r *http.Request) {
	db3 := Openconnection()
	params := mux.Vars(r)
	defer db3.Close()
	rows, err := db3.Query(`SELECT * FROM "newtable" Where "Name"=$1`, params["Name"])
	defer rows.Close()
	var contactl Contact
	for rows.Next() {
		err = rows.Scan(&contactl.Locality, &contactl.Name)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(contactl)

	}
}

func Update(w http.ResponseWriter, r *http.Request) {

	db4 := Openconnection()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	defer db4.Close()
	updateStmt, _ := db4.Prepare(`update "newtable" set "Locality"=$1 where "Name"=$2`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var keyVal map[string]string
	json.Unmarshal(body, &keyVal)
	newLocality := keyVal["Locality"]
	_, err = updateStmt.Exec(newLocality, params["Name"])

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Post with Name = %s was updated", params["Name"])

	fmt.Println("8TH")

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db5 := Openconnection()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, _ := db5.Prepare(`delete from "newtable" where "Name"=$1`)
	fmt.Println("err1")
	_, err := stmt.Exec(params["Name"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("err2")
	fmt.Fprintf(w, "Post with Name = %s was deleted", params["Name"])

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/addcontact", Create).Methods("POST")
	r.HandleFunc("/readcontact", List).Methods("GET")
	r.HandleFunc("/contactbyname/{Name}", Listbyname).Methods("GET")
	r.HandleFunc("/updatecontact/{Name}", Update).Methods("PUT")
	r.HandleFunc("/deletecontact/{Name}", Delete).Methods("DELETE")
	r.HandleFunc("/register", Authentication.Register).Methods("POST")

	http.ListenAndServe(":3008", r)

}
