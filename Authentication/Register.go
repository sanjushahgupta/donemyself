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

	sqlStatement := `INSERT INTO "Jobusers"("id","Name","email","Password","gender")
		VALUES ($1,$2,$3,$4,$5)`
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)

	CreatedUser, err := db1.Exec(sqlStatement, user1.ID, user1.Name, user1.Email, user1.Password, user1.Gender)

	json.NewEncoder(w).Encode(CreatedUser)

}
