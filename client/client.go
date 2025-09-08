package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建HTTP客户端
	client := &http.Client{}

	// 构造请求
	req, err := http.NewRequest("PUT", "http://localhost:8080/good", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Println("状态码:", resp.StatusCode)

	// 如果是空白响应，打印提示
	if resp.ContentLength == 0 {
		fmt.Println("响应是空的")
		return
	}

	// 否则读取并打印响应内容
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println("响应内容:", string(data))
}
