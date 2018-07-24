package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Token []byte `json:"token"`
	jwt.StandardClaims
}

var key = []byte("AllYourBase")

func SetJwt(obj []byte) (tokenString string, err error) {
	// Create the Claims
	var claims = MyCustomClaims{
		obj,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(key)
	return
}
func GetJwt(tokenString string) (tokenByte []byte) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		tokenByte = claims.Token
		fmt.Printf("%v %v", claims.Token, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	return
}
