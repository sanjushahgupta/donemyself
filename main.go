package main

import (
	"firstattemp/Authentication"
	"firstattemp/Crud"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
func main() {

	r := mux.NewRouter()
	r.Use(CommonMiddleware)
	r.HandleFunc("/addcontact", Crud.Create).Methods("POST")
	r.HandleFunc("/readcontact", Crud.List).Methods("GET")
	r.HandleFunc("/contactbyid/{id}", Crud.Listbyid).Methods("GET")
	r.HandleFunc("/updatecontact/{id}", Crud.Update).Methods("PUT")
	r.HandleFunc("/deletecontact/{id}", Crud.Delete).Methods("DELETE")
	r.HandleFunc("/register", Authentication.Register).Methods("POST")
	r.HandleFunc("/login", Authentication.Login).Methods("POST")
	http.ListenAndServe(":1700", r)
}
