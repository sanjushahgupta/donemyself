package Model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Jobdetails struct {
	ID    int    `json:"id"`
	Title string `json:"title"`

	Post       string `json:"post"`
	Salary     int    `json:"salary"`
	Experience string `json:"experience"`
}

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
