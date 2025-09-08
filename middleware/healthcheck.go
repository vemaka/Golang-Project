package middleware

import (
	"Golang/handler"
	"fmt"
	"net/http"
	"time"
)

// 简单健康检查
func HealthCheck(next http.Handler) http.Handler {

	ms := handler.GetRouterConfig()

	// for _, path := range ms.Managerservices {
	// 	fmt.Println("path :", path.Path)
	// }

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			for _, path := range ms.Managerservices {
				url := "http://localhost:8080" + path.Path
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Failed to ping %s: %v\n", url, err)
					continue
				}
				defer resp.Body.Close()
				fmt.Printf("Pinged %s: Status %s\n", url, resp.Status)
			}
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("HealthCheck called")
		next.ServeHTTP(w, req)
	})
}
