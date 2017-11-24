package httpx

import "net/http/cookiejar"

//设置JarCookie容器
func (httpx *Httpx) SetJarCookie(Jar *cookiejar.Jar) {
	httpx.client.Jar = Jar
}

//是否使用JarCookie容器储存cookie (b[true 使用 , false 不使用])
func (httpx *Httpx) SetAutoSaveCookie(b bool) error {
	if b {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		httpx.SetJarCookie(jar)
	}else {
		httpx.SetJarCookie(nil)
	}
	return nil
}