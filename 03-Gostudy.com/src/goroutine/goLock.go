package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg5 sync.WaitGroup
var mutex sync.Mutex
var numtexread sync.RWMutex

func test3() {
	// 添加互斥锁
	mutex.Lock()
	count++
	fmt.Println("this count is ：", count)
	time.Sleep(time.Millisecond)
	// 解互斥锁
	mutex.Unlock()
	wg5.Done()
}

func Write() {
	// 添加互斥锁
	mutex.Lock()
	fmt.Println("执行写操作……")
	time.Sleep(time.Second * 3)
	// 解互斥锁
	mutex.Unlock()
	wg5.Done()
}

func Read() {
	// 添加了读锁
	numtexread.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second * 3)
	// 解开读锁
	numtexread.RUnlock()
	wg5.Done()
}

func main() {
	//for i := 0; i < 20; i++ {
	//	wg5.Add(1)
	//	go test3()
	//}

	for i := 0; i < 20; i++ {
		wg5.Add(1)
		go Write()
	}

	for i := 0; i < 20; i++ {
		wg5.Add(1)
		go Read()
	}

	wg5.Wait()
}
