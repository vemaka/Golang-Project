package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware ：追踪打印日志
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, req)

		duration := time.Since(start)
		if req.Method != "GET" {
			fmt.Printf("| %s | %s | %s | %v | \n", req.Method, req.URL.Path, req.RemoteAddr, duration)
		}

	})
}
