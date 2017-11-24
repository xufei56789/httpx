package httpx

import (
	"net/http"
	"io/ioutil"
)

//整理response数据
func (response *Response) response(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	response.Body = body
	response.Length = resp.ContentLength
	response.Status = resp.StatusCode
	response.Header = resp.Header
	response.Cookies = resp.Cookies()
	response.Url = resp.Request.URL
	return nil
}