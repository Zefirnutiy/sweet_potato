package structs

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	Id     int
	Email  string
	Schema string
}
