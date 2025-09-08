package handler

import (
	r "Golang/router"
	"net/http"
)

// 具体路由配置需要

type Service struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type ManagerService struct {
	Managerservices []Service
}

var serviceList = []Service{
	{"/good", "GET", HeartBeat}, // 健康检查接口
	{"/good", "POST", GOODBuy},
	{"/good", "PUT", GOODSave},
}

// 服务端统一注册所需路由
func RegisterAllServices() {
	for _, service := range serviceList {
		r.RegisterService(service.Path, service.Method, service.Handler)
	}
}

// 获取路由配置
func GetRouterConfig() *ManagerService {
	ms := &ManagerService{
		Managerservices: serviceList,
	}

	return ms
}
