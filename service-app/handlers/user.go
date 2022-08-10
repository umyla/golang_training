package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"service/auth"
	"service/data/user"
	"service/web"
)

type userHandlers struct {
	*user.DbService
	*auth.Auth
}

func (h *userHandlers) SignUp(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return errors.New("value not found in the context")
	}
	var nu user.NewUser
	err := web.Decode(r, &nu)
	if err != nil {
		return err
	}

	//

	usr, err := h.Create(ctx, nu, v.Now)

	if err != nil {
		return fmt.Errorf("user: %+v %w", &usr, err)
	}

	return web.Respond(ctx, w, usr, http.StatusCreated)

}

//func (h *userHandlers)Update(ctx context.Context,w http.Response)

func (h *userHandlers) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return errors.New("value not found in the context")
	}

	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := web.Decode(r, &login)
	if err != nil {
		return err
	}

	claims, err := h.Authenticate(ctx, v.Now, login.Email, login.Password)
	log.Println(err)
	if err != nil {
		return web.NewRequestError(errors.New("please provide a valid email and password"), http.StatusUnauthorized)

	}

	var tkn struct {
		Token string `json:"token"`
	}

	tkn.Token, err = h.GenerateToken(claims)

	if err != nil {
		return fmt.Errorf("generating token %w", err)
	}

	return web.Respond(ctx, w, tkn, http.StatusOK)

}
