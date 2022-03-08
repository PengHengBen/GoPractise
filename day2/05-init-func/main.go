package main

import (
	_ "day2/05-init-func/sub" // 此时，只会调用sub里边的init函数，编译不会出错
	"fmt"
)

func main() {
	//res := sub.Sub(20, 10)
	fmt.Println("sub.Sub(20,10) =")
}
