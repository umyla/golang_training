package handlers

import (
	"context"
	"net/http"
	"service/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "ok",
	}
	//w.Header().Set("Content Type", "application/json")
	//return json.NewEncoder(w).Encode(status)
	return web.Respond(ctx, w, status, http.StatusOK)
}
