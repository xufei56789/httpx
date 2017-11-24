package httpx

import (
	"net/url"
	"net/http"
	"errors"
	"strconv"
	"crypto/x509"
)

//设置代理(proxyUrl [代理地址http://xxxxx:prot])
func (httpx *Httpx) SetProxy(proxyUrl string) error {
	u := url.URL{}
	urlproxy, err := u.Parse(proxyUrl)
	if err != nil {
		return err
	}
	httpx.configTr.Proxy = http.ProxyURL(urlproxy)
	return nil
}

//设置重定向回调函数
func (httpx *Httpx) SetRedirectFunc(f func(req *http.Request, via []*http.Request) (e error)) {
	httpx.client.CheckRedirect = f
}

//重定向次数设定 (i[0 停止重定向, > 0 重定向次数])
func (httpx *Httpx) SetRedirect(i int) {
	//声明闭包
	f := func(req *http.Request, via []*http.Request) (e error) {
		//如果 i = 0 不重定向
		if i <= 0 {
			return http.ErrUseLastResponse
		}
		//如果 当前重定向次数 >= 重定向次数停止
		if len(via) >= i {
			return errors.New("stopped after "+strconv.Itoa(i)+" redirects")
		}
		return nil
	}
	//设置重定向回调函数
	httpx.SetRedirectFunc(f)
}

//是否跳过证书检查
func (httpx *Httpx) SetInsecureSkipVerify(b bool)  {
	httpx.configTls.InsecureSkipVerify = true
}

//设置x509cert证书
func (httpx *Httpx) SetCertPoolx509(b []byte) bool {
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(b)
	if !ok {
		return ok
	}
	httpx.configTls.RootCAs = pool
	return ok
}