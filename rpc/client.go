// rpc/client.go
package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client RPC客户端
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient 创建RPC客户端
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// Call 调用远程方法
func (c *Client) Call(service, method string, args interface{}, result interface{}) error {
	// 准备请求
	argsData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	request := Request{
		Service: service,
		Method:  method,
		Args:    argsData,
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// 发送请求
	resp, err := c.client.Post(c.baseURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode != http.StatusOK {
		return parseError(resp)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if response.Error != "" {
		return &RPCError{Message: response.Error}
	}

	// 解析结果
	resultData, err := json.Marshal(response.Result)
	if err != nil {
		return err
	}

	return json.Unmarshal(resultData, result)
}

// RPCError RPC错误
type RPCError struct {
	Message string
}

type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

func (e *RPCError) Error() string {
	return e.Message
}

func parseError(resp *http.Response) error {
	var errResp struct {
		Error string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return &RPCError{Message: "Unknown error: " + resp.Status}
	}

	return &RPCError{Message: errResp.Error}
}
