package router

import (
	"net/http"
	"regexp"
)

// Router ： 方法的接口展示
type Router interface {
	Get(path string, handler http.HandlerFunc)
	POST(path string, handler http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// router ： 实现路由结构体，将方法 路径 路由规则绑定
type router struct {
	routes map[string]map[string]http.HandlerFunc // method -> path -> handler

	dynamicRoutes map[string]map[*regexp.Regexp]http.HandlerFunc

	notFound http.HandlerFunc
}

// NewRouter ： 实例化对象
func NewRouter() Router {
	return &router{
		routes:        make(map[string]map[string]http.HandlerFunc),
		dynamicRoutes: make(map[string]map[*regexp.Regexp]http.HandlerFunc),
		notFound: func(w http.ResponseWriter, req *http.Request) {
			http.NotFound(w, req)
		},
	}
}

// isDynamicPath ：判断路径是否是动态路由
func isDynamicPath(path string) bool {
	return regexp.MustCompile(`:([a-zA-Z0-9_]+)`).MatchString(path)
}

// parseDynamicRoute ：将动态路由转换为正则表达式
func parseDynamicRoute(path string) *regexp.Regexp {
	// 将动态参数部分（如 :id）替换为正则表达式模式
	pattern := regexp.MustCompile(`:([a-zA-Z0-9_]+)`).ReplaceAllString(path, `([^/]+)`)
	return regexp.MustCompile("^" + pattern + "$")
}

// AddRouter ：添加静态路由
func (r *router) AddRouter(method string, path string, handler http.HandlerFunc) {
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

// AddDynamicRouter ： 添加动态路由
func (r *router) AddDynamicRouter(method string, path string, handler http.HandlerFunc) {
	//将动态路由转换为正则表达式
	pattern := parseDynamicRoute(path)
	if _, ok := r.dynamicRoutes[method]; !ok {
		r.dynamicRoutes[method] = make(map[*regexp.Regexp]http.HandlerFunc)
	}
	r.dynamicRoutes[method][pattern] = handler
}

// TODO : 实现方法 GET
func (r *router) Get(path string, handler http.HandlerFunc) {
	if isDynamicPath(path) {
		r.AddDynamicRouter("GET", path, handler)
	} else {
		r.AddRouter("GET", path, handler)
	}
}

// TODO : 实现方法 POST
func (r *router) POST(path string, handler http.HandlerFunc) {
	if isDynamicPath(path) {
		r.AddDynamicRouter("POST", path, handler)
	} else {
		r.AddRouter("POST", path, handler)
	}
}

// ServeHTTP ：与HTTO路由匹配
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	// 先尝试匹配静态路由
	if handlers, ok := r.routes[method]; ok {
		if handler, found := handlers[path]; found {
			handler(w, req)
			return
		}
	}

	// 如果静态路由没匹配，再尝试动态路由
	if dynamicHandlers, ok := r.dynamicRoutes[method]; ok {
		for pattern, handler := range dynamicHandlers {
			if pattern.MatchString(path) {
				handler(w, req)
				return
			}
		}
	}

	r.notFound(w, req)
}
