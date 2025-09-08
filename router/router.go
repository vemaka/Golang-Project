package router

import (
	"Golang/types"
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

// 获取上下文
func GetHTTPReqFromContext(ctx context.Context) (*types.HTTPReq, error) {

	// ctxValue := ctx.Value("httpReq")
	// fmt.Printf("Context value: %#v, Type: %T\n", ctxValue, ctxValue)

	httpReq, ok := ctx.Value("httpReq").(types.HTTPReq)
	// fmt.Printf("GetHTTPReqFromContext ---------- Context value: %+v\n", httpReq)
	if !ok {
		return nil, errors.New("failed to get HTTPReq from context")
	}
	return &httpReq, nil
}

func (r *Router) AddRouter(path string, method string, handle http.HandlerFunc) {
	r.handlers[path+"|"+method] = RouterHandler{path: path, method: method, handle: handle}
}

// 最后访问路由
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// ctxValue := ctx.Value("httpReq")
	// fmt.Printf("Context value: %+v\n", ctxValue)

	httpReq, err := GetHTTPReqFromContext(ctx)
	if err != nil {
		http.Error(w, "Failed to extract HTTPReq", http.StatusInternalServerError)
		// fmt.Printf("In router ServeHTTP, httpReq: %+v\n", httpReq)
		return
	}

	handle, ok := DiscoverService(httpReq.ServiceName, httpReq.Method)
	if !ok {
		http.NotFound(w, req)
		fmt.Println("router not found")
		return
	}

	handle.ServeHTTP(w, req)
}
