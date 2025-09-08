package router

import (
	"net/http"
)

// 服务发现
func DiscoverService(path string, method string) (http.HandlerFunc, bool) {
	// fmt.Println("DiscoverService ........")
	key := path + "|" + method
	handle, ok := globalRegistry[key]
	return handle, ok
}
