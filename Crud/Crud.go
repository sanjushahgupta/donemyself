package Crud

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	db9 := Dbconnect.Openconnection()
	var data Model.Contact
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("querytest")
	defer db9.Close()
	sqlStatement := `
		INSERT INTO "Jobdetailsdb"("id","Title","Post")
		VALUES ($1,$2,$3)`
	_, err = db9.Exec(sqlStatement, data.ID, data.Title, data.Post)

	if err != nil {
		panic(err)
	}

}

func List(w http.ResponseWriter, r *http.Request) {

	db2 := Dbconnect.Openconnection()
	var newtable Model.Contact
	var Contactsarr []Model.Contact
	defer db2.Close()

	rows, err := db2.Query(`SELECT "id","Title", "Post" FROM "Jobdetailsdb"`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&newtable.ID, &newtable.Title, &newtable.Post)
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

func Listbyid(w http.ResponseWriter, r *http.Request) {
	db3 := Dbconnect.Openconnection()
	params := mux.Vars(r)
	defer db3.Close()
	rows, err := db3.Query(`SELECT * FROM "Jobdetailsdb" Where "id"=$1`, params["id"])
	defer rows.Close()
	var contactl Model.Contact
	for rows.Next() {
		err = rows.Scan(&contactl.ID, &contactl.Title, &contactl.Post)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(contactl)

	}
}

func Update(w http.ResponseWriter, r *http.Request) {

	db4 := Dbconnect.Openconnection()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	defer db4.Close()
	updateStmt, _ := db4.Prepare(`update "Jobdetailsdb" set "Post"=$1 where "id"=$2`)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var keyVal map[string]string
	json.Unmarshal(body, &keyVal)
	newLocality := keyVal["Post"]
	_, err = updateStmt.Exec(newLocality, params["id"])

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Post with id = %s was updated", params["id"])

	fmt.Println("8TH")

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db5 := Dbconnect.Openconnection()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, _ := db5.Prepare(`delete from "Jobdetailsdb" where "id"=$1`)
	fmt.Println("err1")
	_, err := stmt.Exec(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("err2")

}
