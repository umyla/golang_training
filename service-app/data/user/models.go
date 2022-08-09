package user

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Roles        pq.StringArray `json:"roles"`
	PasswordHash []byte         `json:"-"`
	DateCreated  time.Time      `json:"date_created"`
	DateUpdated  time.Time      `json:"date_updated"`
}

type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	Password string   `json:"password"`
}
