package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlLXByb2plY3QiLCJzdWIiOiIxMDIiLCJleHAiOjE2NTk3MDYyODYsImlhdCI6MTY1OTcwMzI4NiwiUm9sZXMiOlsiVVNFUiJdfQ.MM3cx7OogyuJ0a5XLlZO3_THORKMFkQlYVk-e-t7M_b8cA73kJGLPN0r9w8h58pfCLRc2GsxjOzyFYVAcIE3z_h-3LwhF3SzNdoqXlI_EiGVNtQRHiyYU31vBj877UUYK3rWTcRNw9tG2YoYrNtR7tgGMWa99SrVAyKxe65vxiObnvBgKuiVHaFv-FrdS9MUPHtst1vat6hDfVgNzbx2WintwoaA782mwCXmH5k1gLvmKSzDYttC-ply7EPsVBdlr8_QoCoMP7Qz_G9pOQv_lv-kcpOfhhMYHBYolZMj5b0NlI1DoRxB9dNQpAivnuQpEZ3cW6fBstdNmbUVSrLhog`

func main() {
	type claims struct {
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}
	PrivatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatalln("not able to read perm file")
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(PrivatePEM)
	if err != nil {
		log.Fatalln("parsing private key")
	}
	var c claims
	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil

	})
	if err != nil {
		log.Println(err)
	}
	if !token.Valid {
		log.Println("invalid token")
		return
	}
	fmt.Printf("%+v", c)
}
