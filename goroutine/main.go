package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f2(ch chan int) {
	for v := range ch {
		fmt.Println("接收", v)
	}
}

func recv(c chan int) {
	// for v := range c {
	// 	fmt.Println("已接收：", v)
	// 	wg.Done()
	// }
	ret := <-c
	fmt.Println("接收成功", ret)

}
func main() {
	// 无缓存区的
	// ch := make(chan int)
	// // go recv(ch)
	// // wg.Add(1)
	// // ch <- 10
	// // go recv(ch)
	// // wg.Add(1)
	// // ch <- 11
	// // go recv(ch)
	// // wg.Add(1)
	// // ch <- 12
	// // fmt.Println("已发送")
	// // close(ch)
	// // wg.Wait()

	// go f2(ch)
	// fmt.Println("第一次写入")
	// ch <- 10
	// fmt.Println("第二次写入")
	// ch <- 11
	// defer close(ch)
	// time.Sleep(time.Second * 3)

	// 有缓冲的channel
	// ch2 := make(chan int, 3)
	// fmt.Println("缓冲数据个数：", len(ch2), "缓冲区大小：", cap(ch2))

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		fmt.Printf("写子协程：%d,len=%d,cap=%d\n", i, len(ch2), cap(ch2))
	// 		ch2 <- i
	// 	}
	// 	close(ch2)
	// }()

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("读子协程：%d,len=%d,cap=%d\n", i, len(ch2), cap(ch2))
	// 	num := <-ch2
	// 	fmt.Printf("读子协程[%d]已读，读出结果为[%d]: 缓冲区剩余数据个数:%d, 缓冲区大小:%d\n", i, num, len(ch2), cap(ch2))
	// }
	// fmt.Println(ch2)

	// 多路复用
	// 示例中的代码首先是创建了一个缓冲区大小为1的通道 ch，在进入 for 循环后，此时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-c 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
	// 第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；后续的 for 循环同理会依次打印出3、5、7、9。
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Printf("写入%d到channel \n", i)
		}
	}

}
