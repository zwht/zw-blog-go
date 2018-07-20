package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func SetJwt(obj []byte) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, err = token.SignedString(obj)
	return
}
func GetJwt(obj string) {

}
