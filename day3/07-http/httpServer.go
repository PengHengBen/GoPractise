package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 注册路由 router
	// xxxx/user ===> func1
	// xxxx/name ===> func2
	// xxxx/id ===> func3
	// 浏览器输入http://127.0.0.1:8080/user   func是回调函数，用于路由的响应，这个回调函数的原型是固定的
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		// request: ===> 包含客户端发来的数据
		fmt.Println("用户请求详情:")
		fmt.Println("request:", request)

		// writer: ===> 通过writer将数据返回给客户端
		//func WriteString(w Writer, s string) (n int, err error) {}
		_, _ = io.WriteString(writer, "这是/user请求返回的数据!")
	})
	// 浏览器输入http://127.0.0.1:8080/name   func是回调函数，用于路由的响应，这个回调函数的原型是固定的
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name请求返回的数据!")
	})
	// 浏览器输入http://127.0.0.1:8080/id   func是回调函数，用于路由的响应，这个回调函数的原型是固定的
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id请求返回的数据!")
	})

	fmt.Println("Http Server start ...")
	//func ListenAndServe(addr string, handler Handler) error {}
	// 判断错误的简单写法
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed, err:", err)
		return
	}

	//if err != nil{
	//	fmt.Println("http start failed, err:",err)
	//	return
	//}
}
