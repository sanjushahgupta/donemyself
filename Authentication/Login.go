package Authentication

import (
	"encoding/json"
	"firstattemp/Dbconnect"
	"firstattemp/Model"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

func Login(w http.ResponseWriter, r *http.Request) {
	Dbconnect.Openconnection()
	user := &Model.User{}
	fmt.Println(user)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(Email, Password string) map[string]interface{} {
	db1 := Dbconnect.Openconnection()
	user := &Model.User{}
	fmt.Println(user)

	err := db1.Where("Email=?", Email).First(user).Error
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}
	// create token object and add email and standardclaims
	tk := &Model.Token{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// create a new claim with HS256 alogorithm and token claim
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	key := []byte(os.Getenv("SECRET_KEY"))
	fmt.Println(key)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func VerifyToken(r *http.Request) {

	reqToken := r.Header.Get("Authorization")
	key := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil && token.Valid {
		fmt.Println("valid token")

	} else {
		fmt.Println("invalid token")

	}

}
