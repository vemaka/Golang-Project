package main

import (
	"fmt"
	"net/http"

	"Golang_Pro_2/handler"
	"Golang_Pro_2/middleware"
	"Golang_Pro_2/router"
	//ser "Golang_Pro_2/server"
)

func main() {

	http.HandleFunc("/hello", handler.HelloWorld)

	r := router.NewRouter()
	http.Handle("/", r)
	r.Get("/user", http.HandlerFunc(handler.HelloWorld))
	r.Get("/user/:id", http.HandlerFunc(handler.GetOptions))

	// 单个中间件处理
	// stack := middleware.LoggingMiddleware(r)

	// 整合操作
	middlewareStack := middleware.Compose(
		middleware.LoggingMiddleware,
	)

	sr := middlewareStack(r)

	//端口监听服务
	err := http.ListenAndServe(":8080", sr)
	if err != nil {
		fmt.Printf("Failed start server error:&s", err.Error())
	}

}
