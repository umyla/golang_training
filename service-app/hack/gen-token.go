package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {

	PrivatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatalln("not able to read pem file")
	}
	PrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(PrivatePEM)

	if err != nil {
		log.Fatalln("parsing private key")
	}
	claims := struct {
		jwt.RegisteredClaims // fields
		Roles                []string
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "service-project",
			Subject:   "102",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(50 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Roles: []string{"USER"},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	str, err := tkn.SignedString(PrivateKey)

	if err != nil {

		log.Fatalln(err)
	}
	fmt.Println(str)
}
