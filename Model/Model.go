package Model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Jobdetails struct {
	ID         uuid.UUID
	Title      string `json:"title"`
	Post       string `json:"post"`
	Salary     int    `json:"salary"`
	Experience string `json:"experience"`
}

type User struct {
	ID       uuid.UUID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type Token struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	*jwt.StandardClaims
}
type Exception struct {
	Message string `json:"message"`
}
