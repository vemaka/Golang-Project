package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type RouterHandler struct {
	path   string
	method string
	handle http.HandlerFunc
}

type Router struct {
	handlers map[string]RouterHandler
}

func NewRouter() *Router {
	router := &Router{
		handlers: make(map[string]RouterHandler),
	}

	go func() {
		for range updateChan {
			fmt.Println("Updating routers ........")
		}
	}()
	return router
}

type HTTPReq struct {
	ServiceName string
	Method      string
	Body        string
}

func GetHTTPReqFromContext(ctx context.Context) (*HTTPReq, error) {
	httpReq, ok := ctx.Value("httpReq").(HTTPReq)
	if !ok {
		return nil, errors.New("failed to get HTTPReq from context")
	}
	return &httpReq, nil
}

func (r *Router) AddRouter(path string, method string, handle http.HandlerFunc) {
	r.handlers[path+"|"+method] = RouterHandler{path: path, method: method, handle: handle}
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	httpReq, err := GetHTTPReqFromContext(ctx)
	if err != nil {
		http.Error(w, "Failed to extract HTTPReq", http.StatusInternalServerError)
		return
	}

	handle, ok := DiscoverService(httpReq.ServiceName, httpReq.Method)
	if !ok {
		http.NotFound(w, req)
		return
	}

	handle.ServeHTTP(w, req)
}
