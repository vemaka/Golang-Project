package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

// 业务实现内容

// TEST
func GOODSell(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Sell good!!!")
	w.Write([]byte("Sell good!!!!"))
}

// TEST
func GOODSave(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Save good!!!")
	w.Write([]byte("Save good!!!!"))
}

// TEST
func GOODBuy(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Buy good!!!")
	w.Write([]byte("Buy good!!!!"))
}

// 作为健康检查调用
func HeartBeat(w http.ResponseWriter, req *http.Request) {

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
