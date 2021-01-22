package main

import (
	"firstattemp/Authentication"
	"firstattemp/Handle"
	"firstattemp/Middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {

	}
	r := mux.NewRouter()
	r.Use(CommonMiddleware)
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(Middleware.JwtVerify)
	s.HandleFunc("/addjobpost", Handle.Create).Methods("POST")
	r.HandleFunc("/readjobpost", Handle.List).Methods("GET")
	r.HandleFunc("/listjobpostbyid/{id}", Handle.Listbyid).Methods("GET")
	s.HandleFunc("/updatejobpostbyid/{id}", Handle.Update).Methods("PUT")
	s.HandleFunc("/deletejobpost/{id}", Handle.Delete).Methods("DELETE")
	r.HandleFunc("/registeruser", Authentication.Register).Methods("POST")
	r.HandleFunc("/loginuser", Authentication.Login).Methods("POST")
	http.ListenAndServe(":6188", r)

}
