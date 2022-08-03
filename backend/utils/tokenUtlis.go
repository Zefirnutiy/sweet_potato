package utils

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type MyCustomOrganization struct {
	Title string
	jwt.StandardClaims
}

var cfg = db.Load("./settings.cfg")

// Функция нужна для создания токена, при создании организации
func CreateToken(Title string) (string, error) {

	claims := MyCustomOrganization{
		Title,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(12 * time.Hour)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(cfg.SecretWordForOrganisation))

	if err != nil {
		return "", err
	}

	return ss, nil
}

// Функция нужна для создания токена, при создании Клиента
func CreateClientToken(customeStruct jwt.Claims, secretWord string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customeStruct)
	ss, err := token.SignedString([]byte(secretWord))

	if err != nil {
		return "", err
	}

	return ss, nil
}

// Функция нужна для расшифровки токена, для middleware
func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &MyCustomOrganization{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*MyCustomOrganization); ok && token.Valid {
		fmt.Println(claims)
		return claims.Title, nil
	}

	return "", fmt.Errorf("Расшифровки нет")
}
