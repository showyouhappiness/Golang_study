package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		"crf",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}
