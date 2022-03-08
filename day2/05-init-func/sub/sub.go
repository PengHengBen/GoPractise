package sub

import "fmt"

// init函数没有参数，没有返回值，原型固定如下
// init函数是不允许用户显示调用的
// 有时候引用一个包，可能只想使用这个包里边的init函数(mysql的init对驱动进行初始化)
// 但是不想使用这个包的其他函数，为了方式编译器报错，可以使用 _ 形式来处理
// import _ "day2/05-init-func/sub"
func init() {
	fmt.Println("this is first init() in package sub")
}

func init() {
	fmt.Println("this is second init() in package sub")
}

func Sub(a, b int) int {
	//init() // 不允许显示调用
	return a - b
}
