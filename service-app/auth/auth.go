package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type ctxKey int

const Key ctxKey = 1

type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

type Auth struct {
	privateKey *PrivateKey
}

func NewAuth(privateKey *rsa.PrivateKey) (*Auth, error) {
	if privateKey == nil {
		return nil, error.New("privatekey cannot be nil")
	}
	a := Auth{PrivateKey: PrivateKey}
	return &a, nil

}
func (a *Auth) GenerateToken(claims Claims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)
	tokenStr, err := tkn.Signedstring(a.privateKey)
	if err != nil {
		return " ", fmt.Errorf("signes token %w", err)
	}
	return tokenStr, nil
}
func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {

	var claims Claims

	// verifying token with our public key
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.privateKey.Public(), nil
	})

	if err != nil {
		if !token.Valid {
			return Claims{}, errors.New("invalid token")
		}
		return Claims{}, fmt.Errorf("parsing token %w", err)
	}

	return claims, nil

}
