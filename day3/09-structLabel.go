package main

import (
	"encoding/json"
	"fmt"
)

type teacher struct {
	// 反引号中间的字段值一定要用双引号引起来，不然不生效
	Name    string `json:"-"`            // 在使用json编码时，这个字段不参与编码
	Subject string `json:"Subject_name"` // 在json编码时，这个字段会编码成Subject_name
	// 在json编码时，将age转换为string类型，一定要是两个字段紧挨着，字段,类型，中间不能有空格，有空格修改字段的类型不生效
	Age int `json:"age,string"`
	// 在json编码时，如果这个字段是没有赋值，那么忽略掉，不参与编码，如果这个字段已被赋值，参与编码
	Address string `json:"address,omitempty"`

	// 注意，gender是小写的，小写字母开头的，在json编码时会忽略掉
	gender string
}

// 主函数
func main() {
	t1 := teacher{
		Name:    "Duke",
		Subject: "Golang",
		Age:     18,
		gender:  "male",
		Address: "北京",
	}

	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal(&t1)
	fmt.Println("encodeInfo:", string(encodeInfo))
}
