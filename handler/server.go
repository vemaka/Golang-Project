package handler

import (
	r "Golang/router"
	"net/http"
)

type Service struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var serviceList = []Service{
	{"/good", "GET", GOODSell},
	{"/good", "POST", GOODBuy},
	{"/good", "PUT", GOODSave},
}

func RegisterAllServices() {
	for _, service := range serviceList {
		r.RegisterService(service.Path, service.Method, service.Handler)
	}
}
