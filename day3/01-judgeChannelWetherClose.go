package main

import "fmt"

func main() {
	numChan := make(chan int, 10)

	go func() {
		for i := 0; i < 50; i++ {
			numChan <- i
			fmt.Println("写入数据:", i)
		}
		close(numChan)
	}()

	// read
	//for v := range numChan {
	//	fmt.Println("v:",v)
	//}

	for {
		v, ok := <-numChan // ok-idom模式判断
		if !ok {
			fmt.Println("管道已经关闭，准备退出")
			break
		}
		fmt.Println("v", v)
	}

	fmt.Println("OVER!")
}
