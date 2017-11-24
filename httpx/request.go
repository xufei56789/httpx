package httpx

import (
	"net/http"
	"bytes"
)

//发起http请求
func (httpx *Httpx) Request(method string, url string, b []byte) (*Response, error) {
	//创建请求
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != err {
		return nil, err
	}
	//判断是否存在header指针
	if httpx.header != nil {
		req.Header = *httpx.header
	}
	resp, err := httpx.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	err = response.response(resp)
	if err != nil {
		return response, err
	}
	return response, error(nil)
}