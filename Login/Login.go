package Login

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &Model.User{}
	// Read the content of the request body.
	// Cast the request body content into our user struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
	}
	// passing the email and password inside the findone function and invoke it
	resp := Findone(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)

}

func Findone(Email, Password string) map[string]interface{} {
	db1 := Dbconnect.Openconnection()
	user := &Model.User{}
	Sqlstmt := `SELECT *FROM users WHERE email=$1`
	_, err := db1.Exec(Sqlstmt, user.Email)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email Address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		// password doesnot match
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials.Please try again"}
		return resp
	}

	tk := &Model.Token{
		UserID: user.Id,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}
