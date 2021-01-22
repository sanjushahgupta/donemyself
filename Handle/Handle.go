package Handle

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"io/ioutil"
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
	db := Dbconnect.Openconnection()
	var contactarr []Model.Jobdetails
	params := mux.Vars(r)["id"]
	defer db.Close()
	db.Where("id = $1", params).First(&contactarr)
	json.NewEncoder(w).Encode(contactarr)
}

// func Update(w http.ResponseWriter, r *http.Request) {
// 	db := Dbconnect.Openconnection()
// 	var data []Model.Jobdetails
// 	var sse Model.Jobdetails
// 	params := mux.Vars(r)["id"]
// 	db.Where("id = $1", params).Find(&data)
// 	ss := Model.Jobdetails{Title: sse.Title, Post: sse.Post, Salary: sse.Salary, Experience: sse.Experience}

// 	db.Update(&ss)
// 	json.NewEncoder(w).Encode(ss)
// }

func Update(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent Model.Jobdetails
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range Model.Jobdetails {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Post = updatedEvent.Post
			singleEvent.Salary = updatedEvent.Salary
			singleEvent.Experience = updatedEvent.Experience
			events := append(events[:i], singleEvent)
			fmt.Println(events)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := Dbconnect.Openconnection()
	var contactarr []Model.Jobdetails
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	db.Where("id = $1", params).Delete(&contactarr)

}
