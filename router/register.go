package router

import (
	"net/http"
)

var globalRegistry = make(map[string]http.HandlerFunc)
var updateChan = make(chan struct{})

// 服务注册
func RegisterService(path string, method string, handle http.HandlerFunc) {
	// fmt.Println("RegisterService ........")
	key := path + "|" + method
	globalRegistry[key] = handle
	updateChan <- struct{}{}
	NewRouter().AddRouter(path, method, handle)
}
