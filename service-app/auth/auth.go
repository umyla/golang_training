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

const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

func (c Claims) HasRoles(roles ...string) bool {

	for _, has := range c.Roles { // roles with the user
		for _, want := range roles { // what roles  handler demand
			if has == want {
				return true
			}
		}
	}

	return false

}

type Auth struct {
	privateKey *rsa.PrivateKey
}

func NewAuth(privateKey *rsa.PrivateKey) (*Auth, error) {
	if privateKey == nil {
		return nil, errors.New("privatekey cannot be nil")
	}
	a := Auth{privateKey: privateKey}
	return &a, nil

}
func (a *Auth) GenerateToken(claims Claims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)
	tokenStr, err := tkn.SignedString(a.privateKey)
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
