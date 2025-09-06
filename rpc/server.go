package rpc

import (
	"encoding/json"
	"net/http"
	"reflect"
)

// Service: 定义RPC服务接口
type Service interface {
	Name() string
}

// Server: 定义服务器结构体
type Server struct {
	services map[string]Service
}

// NewServer： 创建新服务器
func NewServer() *Server {
	return &Server{
		services: make(map[string]Service),
	}
}

// RegisterService： 注册服务器
func (s *Server) RegisterService(service Service) {
	s.services[service.Name()] = service
}

// POSTHandlerRPC： RPC的POST方法
func (s *Server) POSTHandlerRPC(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Method not allowed"))
		return
	}

	// 解析请求
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	// 查找服务
	service, exist := s.services[req.Service]
	if !exist {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// 使用反射调用方法
	serviceValue := reflect.ValueOf(service)
	method := serviceValue.MethodByName(req.Method)
	if !method.IsValid() {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	// 准备参数
	args := reflect.New(method.Type().In(0).Elem())
	if err := json.Unmarshal(req.Args, args.Interface()); err != nil {
		http.Error(w, "Invalid arguments", http.StatusNotFound)
		return
	}

	// 调用方法
	result := method.Call([]reflect.Value{args})

	// 处理结果
	if err, ok := result[1].Interface().(error); ok && err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 返回响应
	response := Respons{
		Result: result[0].Interface(),
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

type Request struct {
	Service string          `json:"service"`
	Method  string          `json:"method"`
	Args    json.RawMessage `json:"args"`
}

type Respons struct {
	Result interface{} `json"result"`
	Erroc  string      `json:"error,omitempty"`
}
