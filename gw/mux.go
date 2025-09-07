package gw

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	r "Golang/router"
	"Golang/types"
)

type gw_mux struct {
	s      *Server
	router *r.Router
}

type Server struct {
	Name string
	Addr string
	Mux  http.ServeMux
}

func NewGWMux() *gw_mux {
	mux := &gw_mux{
		s: &Server{
			Mux: *http.NewServeMux(),
		},
		router: r.NewRouter(),
	}

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

	// fmt.Println("Entering gw_mux ServeHTTP")

	str, err := extract_parameter(req)
	if err != nil {
		fmt.Println("Get json name failed")
		return
	}

	// 组装请求
	httpReq := types.HTTPReq{
		ServiceName: req.URL.Path,
		Method:      req.Method,
		Body:        str,
	}

	// 调用下游服务接口
	ctx := context.WithValue(req.Context(), "httpReq", httpReq)
	req = req.WithContext(ctx)

	// fmt.Printf("httpReq: %+v\n", httpReq)
	// r.RegisterService("/good","GET",handle.GOODBuy)
	// r.RegisterService("/good","POST",handle.GOODSell)
	// r.RegisterService("/good","PUT",handle.GOODSave)

	m.router.ServeHTTP(w, req)
}
