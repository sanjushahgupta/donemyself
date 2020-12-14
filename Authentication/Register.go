package Authentication

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
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

func Register(w http.ResponseWriter, r *http.Request) {
	// Read the content of the request body.
	db1 := Openconnection()
	user1 := &User{}

	json.NewDecoder(r.Body).Decode(user1)

	// Encrypt the user password before saving.
	pass, err := bcrypt.GenerateFromPassword([]byte(user1.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(err)

	// Save the user information.
	user1.Password = string(pass)

	sqlStatement := `INSERT INTO "users"("Name","email","Password","gender")
		VALUES ($1,$2,$3,$4)`
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)

	CreatedUser, err := db1.Exec(sqlStatement, user1.Name, user1.Email, user1.Password, user1.Gender)

	json.NewEncoder(w).Encode(CreatedUser)

}
