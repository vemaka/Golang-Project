package router

import "net/http"

var globalRegistry = make(map[string]http.HandlerFunc)
var updateChan = make(chan struct{})

func RegisterService(path string, method string, handle http.HandlerFunc) {
    key := path + "|" + method
    globalRegistry[key] = handle
    updateChan <- struct{}{}
}