package structs

import "github.com/dgrijalva/jwt-go/v4"

type OrganizationClaims struct {
	jwt.StandardClaims
	Id int32
}
