package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	Token []byte `json:"token"`
	jwt.StandardClaims
}

var key = []byte("AllYourBase")

func SetJwt(obj []byte) (tokenString string, err error) {
	// Create the Claims
	claims := MyCustomClaims{
		obj,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(key)
	return
}
func GetJwt(tokenString string) (pass bool, tokenByte []byte) {
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			pass = false
		}
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			tokenByte = claims.Token
			pass = true
			//fmt.Printf("%v %v", claims.Token, claims.StandardClaims.ExpiresAt)
		} else {
			pass = false
		}
	} else {
		pass = false
	}
	return
}
