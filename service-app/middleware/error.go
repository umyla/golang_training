package middleware

import (
	"context"
	"errors"
	"net/http"
	"service/web"
)

func (m *Mid) Error(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		v, ok := ctx.Value(web.KeyValue).(*web.Values)
		if !ok {
			return errors.New("value not found in the context")
		}
		err := next(ctx, w, r)
		if err != nil {
			m.Log.Printf("%s:Error    :%v", v.TraceId, err)
			err = web.RespondError(ctx, w, err)
			if err != nil {

				return err
			}

		}
		return nil
	}
}
