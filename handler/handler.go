package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

// TEST
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello World!!!")
	w.Write([]byte("Hello World!!!!"))
}

// GetOptions : 实现Get方法的具体内容
func GetOptions(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	re := regexp.MustCompile(`^/user/([^/]+)$`)
	matches := re.FindStringSubmatch(path)

	if len(matches) > 1 {
		id := matches[1]
		fmt.Printf("Found user with ID : %s", id)
	} else {
		http.NotFound(w, req)
	}
}


