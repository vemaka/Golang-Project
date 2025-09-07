package main

import (
	// handlers "Golang/handler"
	"fmt"
	"net/http"
)

// 注册统一API接口
func main() {

	g := NewGWMux()

	mux := http.NewServeMux()

	mux.Handle("/",g)
	

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("8080 port listen failed")
	}

}
