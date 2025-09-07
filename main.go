package main

import (
	gw "Golang/gw"
	"Golang/handler"
	"fmt"
	"net/http"
)

// 注册统一API接口
func main() {

	g := gw.NewGWMux()

	mux := http.NewServeMux()

	mux.Handle("/", g)

	handler.RegisterAllServices()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("8080 port listen failed")
	}

}
