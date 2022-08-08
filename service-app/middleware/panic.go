package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"service/web"
)

func (m *Mid) Panic(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
		v, ok := ctx.Value(web.KeyValue).(*web.Values)
		if !ok {
			return errors.New("value not found")
		}
		defer func() {
			r := recover()
			if r != nil {
				err = fmt.Errorf("PANIC : %v", r)
				log.Printf("%s: \n%s", v.TraceId, debug.Stack())
			}
		}()
		return next(ctx, w, r)
	}
}
