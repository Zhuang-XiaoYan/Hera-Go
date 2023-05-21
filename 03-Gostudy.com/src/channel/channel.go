package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要制定通道中的元素的类型

// 通道必须使用make函数来才能初始化 才能使用
// 通道共有
// 发送（send）  ch <- 10 // 把10发送到ch中
// 接收(receive）  x := <- ch // 从ch中接收值并赋值给变量x   <-ch // 从ch中接收值，忽略结果
// 关闭（close）  close(ch)

var wg sync.WaitGroup
var once sync.Once
var b1 chan int

func noBuff() {
	fmt.Println(b)
	b = make(chan int) // 无缓冲区
	fmt.Println("-----------------------------------")
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := <-b
		fmt.Println("后台goroutine从通道中读取到了", res) // 后台goroutine 从通道中的初始化
	}()
	b <- 10 // hangzhule
	fmt.Println("-----------------------------------")
	wg.Wait()
}

func withBuff() {
	fmt.Println(b)
	b = make(chan int, 16) // 有缓冲区
	fmt.Println("-----------------------------------")
	b <- 10 // hangzhule
	fmt.Println("-----------------------------------")
}

func task1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func task2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) })
}

func main() {
	// withBuff()
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go task1(a)
	go task2(a, b)
	go task2(a, b)
	wg.Wait()
	for res := range b {
		fmt.Println(res)
	}

	fmt.Println("---------------------------------------------")
	chan2 := make(chan<- int, 2) // 表示只能写的管道
	chan3 := make(<-chan int, 3) // 表示只能读的管道

	fmt.Println(chan2)
	fmt.Println(chan3)

}
