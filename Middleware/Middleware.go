package Middleware

import (
	"context"
	"encoding/json"
	"firstattemp/Model"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := &Model.User{}
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		//Grab the token from the header

		headers := strings.TrimSpace(header)

		if headers == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Model.Exception{Message: "Missing auth token"})
			return
		}
		expiresAt := time.Now().Add(time.Minute * 100000).Unix()
		tk := &Model.Token{
			Name:  user.Name,
			Email: user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		_, err := jwt.ParseWithClaims(headers, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Model.Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
