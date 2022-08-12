package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"service/auth"
	"service/data/user"
	"service/web"
)

func (h *userHandlers) AddInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return errors.New("value not found in the context")
	}
	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return web.NewRequestError(errors.New("not authenticated"), http.StatusUnauthorized)
	}
	var newInv user.NewShirtInventory

	err := web.Decode(r, &newInv)
	if err != nil {
		return nil
	}
	s, err := h.CreateInventory(ctx, newInv, claims.Subject, v.Now)
	if err != nil {
		log.Println(err)
		return web.NewRequestError(errors.New("problem in creating new error"), http.StatusBadRequest)
	}
	return web.Respond(ctx, w, s, http.StatusCreated)
}
func (h *userHandlers) ViewInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return web.NewRequestError(errors.New("not authenticated"), http.StatusUnauthorized)
	}
	shirts, err := h.ViewAll(ctx, claims.Subject)
	if err != nil {
		return web.NewRequestError(errors.New("problem in viewing inventory"), http.StatusBadRequest)
	}
	return web.Respond(ctx, w, shirts, http.StatusFound)
}
