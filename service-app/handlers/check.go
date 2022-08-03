package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "ok",
	}
	w.Header().Set("Content Type", "Application/json")
	return json.NewEncoder(w).Encode(status)
}
