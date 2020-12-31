package Authentication

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Read the content of the request body.
	db1 := Dbconnect.Openconnection()
	var user1 Model.User
	ids := uuid.New()

	json.NewDecoder(r.Body).Decode(&user1)

	// Encrypt the user password before saving.
	pass, err := bcrypt.GenerateFromPassword([]byte(user1.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(err)

	// Save the user information.
	user1.Password = string(pass)
	CREATE := Model.User{ID: ids, Name: user1.Name, Gender: user1.Gender, Password: user1.Password, Email: user1.Email}
	createdUser := db1.Create(&CREATE)
	json.NewEncoder(w).Encode(createdUser)

}
