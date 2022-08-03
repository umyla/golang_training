package model

import (
	"errors"
)

type user struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
}

var ErrUserNotFound = errors.New("user id is not found")
var users = map[uint64]user{
	123: {
		FName: "Sam",
		LName: "naik",
		Email: "samnaik@email.com",
	},
	124: {
		FName: "jack",
		LName: "wills",
		Email: "jackwill@email.com",
	},
}

func FetchUser(id uint64) (user, error) {
	u, ok := users[id]
	if !ok {
		return user{}, ErrUserNotFound
	}
	return u, nil
}
