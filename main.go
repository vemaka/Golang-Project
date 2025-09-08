package main

import (
	gw "Golang/gw"
	"Golang/handler"
	"Golang/middleware"
	"fmt"
	"net/http"
)

// 注册统一API接口
func main() {

	g := gw.NewGWMux()

	mux := http.NewServeMux()

	mux.Handle("/", g)

	handler.RegisterAllServices()

	// 中间件管理/应用
	manager := middleware.NewMiddlewareManager()
	manager.Use(middleware.LoggingMiddleware)
	manager.Use(middleware.HealthCheck)

	handle := manager.Apply(mux)

	err := http.ListenAndServe(":8080", handle)
	if err != nil {
		fmt.Println("8080 port listen failed")
	}

}
