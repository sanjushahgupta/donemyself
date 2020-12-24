package Authentication

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Read the content of the request body.
	db1 := Dbconnect.Openconnection()
	user1 := &Model.User{}

	json.NewDecoder(r.Body).Decode(user1)

	// Encrypt the user password before saving.
	pass, err := bcrypt.GenerateFromPassword([]byte(user1.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(err)

	// Save the user information.
	user1.Password = string(pass)
	createdUser := db1.Create(user1)
	if createdUser != nil {
		fmt.Println("error is")
	}
	json.NewEncoder(w).Encode(createdUser)

}
