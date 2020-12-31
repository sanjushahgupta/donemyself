package Handle

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {

	db := Dbconnect.Openconnection()
	var data Model.Jobdetails
	errs := json.NewDecoder(r.Body).Decode(&data)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	ids := uuid.New()

	contacts := Model.Jobdetails{ID: ids, Title: data.Title, Post: data.Post, Salary: data.Salary, Experience: data.Experience}
	db.Create(&contacts)

	fmt.Println("created")
	defer db.Close()

}

func List(w http.ResponseWriter, r *http.Request) {
	var contactarr []Model.Jobdetails

	db := Dbconnect.Openconnection()
	db.Find(&contactarr)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contactarr)

}

func Listbyid(w http.ResponseWriter, r *http.Request) {
	var contactarr []Model.Jobdetails
	db := Dbconnect.Openconnection()
	params := mux.Vars(r)["title"]
	defer db.Close()
	db.Where("title = $1", params).First(&contactarr)
	json.NewEncoder(w).Encode(contactarr)

}

/*func Update(w http.ResponseWriter, r *http.Request) {
	db := Dbconnect.Openconnection()
	// var data Model.Jobdetails
	var contactarr []Model.Jobdetails
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	fmt.Println(params)
	defer db.Close()

	db.Model(&contactarr).Where("id =$1", params).Update()

	// Update("post", "studio-producer")

}*/

func Delete(w http.ResponseWriter, r *http.Request) {
	db := Dbconnect.Openconnection()
	var contactarr []Model.Jobdetails

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["title"]
	db.Where("title = $1", params).Delete(&contactarr)

}
