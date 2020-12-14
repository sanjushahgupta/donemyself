package Model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type Token struct {
	UserID int
	Name   string
	Email  string
	*jwt.StandardClaims
}
