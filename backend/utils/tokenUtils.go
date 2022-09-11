package utils

import (
	"fmt"
	"time"

	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/dgrijalva/jwt-go/v4"
	"os"
)

type MyCustomOrganization struct {
	Organization structs.Claims
	jwt.StandardClaims
}


// Функция нужна для создания токена, при создании организации
func CreateToken(Model structs.Claims) (string, error) {

	claims := MyCustomOrganization{
		Model,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(12 * time.Hour)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("SECRET_WORD_FOR_ORG")))

	if err != nil {
		return "", err
	}

	return ss, nil

}

// Функция нужна для расшифровки токена, для middleware
func ParseToken(accessToken string, signingKey []byte) (structs.Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &MyCustomOrganization{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return structs.Claims{}, err
	}

	if claims, ok := token.Claims.(*MyCustomOrganization); ok && token.Valid {
		return claims.Organization, nil
	}

	return structs.Claims{}, fmt.Errorf("расшифровки нет")
}
