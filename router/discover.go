package router

import (
	"net/http"
)

func DiscoverService(path string, method string) (http.HandlerFunc, bool) {
	// fmt.Println("DiscoverService ........")
	key := path + "|" + method
	handle, ok := globalRegistry[key]
	return handle, ok
}
