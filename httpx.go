package httpx

import (
	"net/http"
	"net/url"
	"crypto/tls"
)

type Httpx struct {
	//http客户端指针
	client *http.Client
	//协议头指针
	header *http.Header
	//配置指针
	configTr *http.Transport
	configTls *tls.Config
}

type Response struct {
	//Body
	Body []byte
	//Length
	Length int64
	//status code
	Status int
	//Headers
	Header http.Header
	//Cookies
	Cookies []*http.Cookie
	//URL
	Url *url.URL
}

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	TRACE   = "TRACE"
	CONNECT = "CONNECT"
)

//创建httpx指针对象
func NewHttpx() *Httpx {
	var httpx Httpx
	//声明配置 
	httpx.client = &http.Client{}
	return &httpx
}