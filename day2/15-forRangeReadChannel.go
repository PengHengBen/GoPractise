package main

import (
	"fmt"
	"time"
)

/*
管道总结：
1. 当管道写满了，写阻塞
2. 当缓冲区读完了，读阻塞
3. 如果管道没有使用make分配空间，管道默认是nil，从nil的管道读取数据、写入数据，都会阻塞（注意：不会崩溃）
4. 从一个已经close的管道读取数据时，会返回零值（不会崩溃）
5. 向一个已经close的管道写入数据时，会崩溃
6. 关闭一个已经close的管道，程序会崩溃
7. 关闭管道的动作，必须在写端，不能在读端，否则写端继续写会崩溃
----> ?
8. 读和写的次数，已经要对等，否则：
    1> 在多个go程中：资源泄露
    2> 在主go程中，程序崩溃（deadlock）
 */

func main() {
	numsChan2 := make(chan int,10)

	// 写
	go func() {
		for i:=0; i<50; i++ {
			numsChan2 <- i
			fmt.Println("写入数据:",i)
		}
		fmt.Println("数据全部写入完毕，准备关闭管道!")
		close(numsChan2) // 如果不关闭管道，for range遍历时，会一直等待，造成死锁，程序崩溃
		//numsChan2 <- 10 // panic: send on closed channel 向一个关闭的管道写数据时，程序会崩溃
		//close(numsChan2) //panic: close of closed channel 关闭一个已经关闭的管道，程序会崩溃
	}()

	// 遍历管道时，只返回一个值
	// for range是不知道管道是否已经写完了，所以只会一直在这里等待
	// 在写入端，将管道关闭，for range遍历关闭的管道时，会退出
	for v := range numsChan2{
		fmt.Println("读取数据:",v)
	}

	time.Sleep(3*time.Second)
	i := <-numsChan2
	fmt.Println("已经关闭之后，继续读取:",i) // 返回0值，不会崩溃

	fmt.Println("OVER!")
}
