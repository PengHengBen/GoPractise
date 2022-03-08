package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http包
	client := http.Client{}

	//func (c *Client) Get(url string) (resp *Response, err error) {}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Client.Get err:", err)
		return
	}

	// 放在上面，内容很多
	body := resp.Body
	// 直接打印，拿不到内容，需要解析 body 111: &{0xc00004c340 {0 0} false <nil> 0xa93500 0xa93480}
	fmt.Println("body 111:", body)
	// func ReadAll(r io.Reader) ([]byte, error) {}
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("body string:", string(readBodyStr))

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("content-type:", ct)
	fmt.Println("date:", date)
	// BWS/1.1 BWS是Baidu Web Server，是百度开发的一个web服务器，大部分百度的web应用程序使用的是BWS
	fmt.Println("server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	fmt.Println("url:", url)       // https://www.baidu.com
	fmt.Println("code:", code)     // 200
	fmt.Println("status:", status) // 200 ok
}
