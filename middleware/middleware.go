package middleware

import (
	"net/http"
)

// Middleware ：
type Middleware func(http.Handler) http.Handler

// Compose ： 整合中间件
func Compose(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {

		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
