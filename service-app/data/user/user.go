package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"service/auth"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var ErrAuthenticationFailure = errors.New("authentication failed")

type DbService struct {
	*sql.DB
}

func (db *DbService) Create(ctx context.Context, nu NewUser, now time.Time) (User, error) {

	// converting the plain text password into the hash using the bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash %w", err)
	}

	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: hash,
		DateCreated:  now,
		DateUpdated:  now,
	}

	const q = `INSERT INTO users
		(name, email, password_hash, roles, date_created, date_updated)
		VALUES ( $1, $2, $3, $4, $5, $6)
		Returning id`

	var id int

	row := db.QueryRowContext(ctx, q, u.Name, u.Email, u.PasswordHash, u.Roles, u.DateCreated, u.DateUpdated)
	err = row.Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("inserting user %w", err)
	}

	u.ID = strconv.Itoa(id)

	return u, nil

}

func (db DbService) Authenticate(ctx context.Context, now time.Time, email, password string) (auth.Claims, error) {

	const q = `SELECT id,name,email,roles,password_hash FROM users WHERE email = $1`

	var u User
	err := db.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.Name, &u.Email, &u.Roles, &u.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return auth.Claims{}, ErrAuthenticationFailure
		}
	}

	err = bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return auth.Claims{}, ErrAuthenticationFailure
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "service project",
			Subject:   u.ID,
			Audience:  jwt.ClaimStrings{"students"},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Roles: u.Roles,
	}

	return claims, nil

}
