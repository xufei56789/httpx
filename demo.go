package main

import (
	"httpx/httpx"
	"fmt"
)

func main()  {
	hx := httpx.NewHttpx()
	//跳过证书验证
	hx.SetInsecureSkipVerify(true)
	//启用cookie容器储存cookie
	hx.SetAutoSaveCookie(true)
	//设置302跳转 0为不跳转
	hx.SetRedirect(0)
	resp , err := hx.Get("https://kyfw.12306.cn/otn/login/init")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resp.Body))
}