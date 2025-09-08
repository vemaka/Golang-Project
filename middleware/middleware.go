package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// MiddlewareManager : middleware 管理器 
type MiddlewareManager struct {
	middlewares []Middleware
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{
		middlewares: make([]Middleware, 0),
	}
}

func (mm *MiddlewareManager) Use(middleware Middleware) {
	mm.middlewares = append(mm.middlewares, middleware)
}

func (mm *MiddlewareManager) Apply(handler http.Handler) http.Handler {
	for i := len(mm.middlewares) - 1; i >= 0; i-- {
		handler = mm.middlewares[i](handler)
	}
	return handler
}
