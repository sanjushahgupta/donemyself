package main

import (
	"firstattemp/Handle"
	"firstattemp/Struct"
	"net/http"

	"github.com/gorilla/mux"
)

// // Lets start with the main function,
// firstly we initialize the router variable as r
// and the GorillaMux is used by calling mux.NewRouter().
// Then we add all the HandleFunc() methods
// which a basic CRUD application will have.
// Here, I have used some named-parameters {name},
// so that we can access the contact details using the name of the person.
func main() {
	Handle.Persons = append(Handle.Persons, Struct.Contact{Name: "sanjuss", StreetNo: 67, Locality: "Kup"})
	Handle.Persons = append(Handle.Persons, Struct.Contact{Name: "sushil", StreetNo: 98, Locality: "Baneswor"})
	r := mux.NewRouter()

	r.HandleFunc("/addcontact", Handle.Create).Methods("POST")
	r.HandleFunc("/readcontact", Handle.Read).Methods("GET")
	r.HandleFunc("/updatecontact/{name}", Handle.Update).Methods("PUT")
	r.HandleFunc("/deletecontact/{name}", Handle.Delete).Methods("DELETE")
	http.ListenAndServe(":8001", r)
}
