package web

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	v, ok := ctx.Value(KeyValue).(*Values)
	if !ok {
		return errors.New("value not found in the context")
	}
	v.StatusCode = statusCode
	w.Header().Set("Content Type", "Application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {
	var webErr *Error
	ok := errors.As(err, &webErr)
	if ok {
		err := webErr.Err
		er := ErrorResponse{Error: err.Error()}
		err = Respond(ctx, w, er, webErr.Status)
		if err != nil {
			return err
		}
		return nil
	}
	er := ErrorResponse{Error: http.StatusText(http.StatusInternalServerError)}
	err = Respond(ctx, w, er, http.StatusInternalServerError)
	if err != nil {
		return err
	}
	return nil
}
