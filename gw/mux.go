package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type gw_mux struct {
	s *Server
}

type Server struct {
	Name string
	Addr string
	Mux  http.ServeMux
}

type HTTPReq struct {
	ServiceName string
	Method      string
	Body        string
}

func NewGWMux() *gw_mux {
	mux := new(gw_mux)

	return mux
}

// return json string
func extract_parameter(req *http.Request) (string, error) {

	var ret string

	if req.Method == "GET" {

		v := req.URL.Query()

		vv := make(map[string]string, len(v))

		for k, item := range v {
			vv[k] = item[0]
		}

		b, err := json.Marshal(vv)

		if err != nil {
			return "", err
		}

		ret = string(b)

	}

	if req.Method == "POST" {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return "", err
		}

		ret = string(b)
	}

	fmt.Println(ret)

	return ret, nil

}

func (m *gw_mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	str,err := extract_parameter(req)
	if err != nil {
		fmt.Println("Get json name failed")
		return
	}

	// 组装请求
	httpReq := HTTPReq{
		ServiceName: req.URL.Path,
		Method: req.Method,
		Body: str,
	}

	// 调用下游服务接口
	ctx := context.WithValue(req.Context(),"httpReq",httpReq)
	req = req.WithContext(ctx)

	m.s.Mux.ServeHTTP(w, req)
}
