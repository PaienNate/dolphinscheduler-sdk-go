package remote

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
)

// DolphinsRestTemplate 提供自定义的 REST 方法，使用 grequests 库。
type DolphinsRestTemplate struct {
	client *grequests.RequestOptions
}

// NewDolphinsRestTemplate 初始化一个新的 DolphinsRestTemplate。
func NewDolphinsRestTemplate(headers map[string]string) *DolphinsRestTemplate {
	return &DolphinsRestTemplate{
		client: &grequests.RequestOptions{
			Headers: headers, // 设置请求头部信息
		},
	}
}

// Get 执行 HTTP GET 请求。
func (drt *DolphinsRestTemplate) Get(url string, query map[string]string) (*HttpRestResult, error) {
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Params:  query,              // 设置查询参数
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// PostForm 执行带表单数据的 HTTP POST 请求。
func (drt *DolphinsRestTemplate) PostForm(url string, body map[string]string) (*HttpRestResult, error) {
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
		Data:    body,               // 设置请求体为表单数据
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// PostJson 执行带 JSON 数据的 HTTP POST 请求。
func (drt *DolphinsRestTemplate) PostJson(url string, body interface{}) (*HttpRestResult, error) {
	data, err := json.Marshal(body) // 将请求体序列化为 JSON
	if err != nil {
		return nil, err
	}
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
		JSON:    data,               // 设置请求体为 JSON 数据
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// PutForm 执行带表单数据的 HTTP PUT 请求。
func (drt *DolphinsRestTemplate) PutForm(url string, body map[string]string) (*HttpRestResult, error) {
	resp, err := grequests.Put(url, &grequests.RequestOptions{
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
		Data:    body,               // 设置请求体为表单数据
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// PutJson 执行带 JSON 数据的 HTTP PUT 请求。
func (drt *DolphinsRestTemplate) PutJson(url string, body interface{}) (*HttpRestResult, error) {
	data, err := json.Marshal(body) // 将请求体序列化为 JSON
	if err != nil {
		return nil, err
	}
	resp, err := grequests.Put(url, &grequests.RequestOptions{
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
		JSON:    data,               // 设置请求体为 JSON 数据
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// Delete 执行 HTTP DELETE 请求。
func (drt *DolphinsRestTemplate) Delete(url string, query map[string]string) (*HttpRestResult, error) {
	resp, err := grequests.Delete(url, &grequests.RequestOptions{
		Headers: drt.client.Headers, // 使用初始化时的请求头部信息
		Params:  query,              // 设置查询参数
	})
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

// handleResponse 处理 grequests 的响应。
func handleResponse(resp *grequests.Response) (*HttpRestResult, error) {
	if !resp.Ok {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode) // 如果响应状态码不为 200，则返回错误
	}

	result := &HttpRestResult{}
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// HttpRestResult 表示 HTTP 请求的结果。
// 转换而来
type HttpRestResult struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data"` // 响应数据
	Failed  bool            `json:"failed"`
	Success bool            `json:"success"`
}
