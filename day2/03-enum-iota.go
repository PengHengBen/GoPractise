package main

import "fmt"

// 在go语言中没有枚举类型，但是我们可以使用const + iota(常量累加器)来进行模拟
// 模拟一周的枚举
const (
	MONDAY    = iota       // 0
	TUESDAY   = iota       // 1
	WEDNESDAY = iota       // 2
	THURSDAY               // 3 ==> 没有赋值，默认与上一行相同 iota ==> 3
	FRIDAY                 // 4
	SATURDAY               // 5
	SUNDAY                 // 6
	M, N      = iota, iota // const属于预编译期赋值，所以不需要:=进行自动推导
)

const (
	JAN  = iota + 1 // 1
	FEB             // 2
	MAR             // 3
	APRI            // 4
)

// 1. iota是常量组计数器
// 2. iota从0开始，每换行递增1
// 3. 常量组如果不赋值，默认与上一行表达式相同
// 4. 如果同一行出现两个iota，那么两个iota的值是相同的
// 5. 每个常量组的iota是独立的，如果遇到const iota会重新清零

// main函数
func main() {

	// 可以使用变量组来统一定义变量
	//var (
	//	number int
	//	name   string
	//	flag   bool
	//)
	fmt.Println("打印周：")
	fmt.Println(MONDAY, TUESDAY, WEDNESDAY, THURSDAY)
	fmt.Println("M:",M, "N:",N)
	fmt.Println("打印月：")
	fmt.Println(JAN,FEB,MAR,APRI)
}
