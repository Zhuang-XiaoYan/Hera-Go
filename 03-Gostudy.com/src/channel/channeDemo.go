package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

// 向管道存储数据
func PutNum(intchan chan int) {
	defer wg1.Done()
	for i := 2; i < 120000; i++ {
		intchan <- i
	}
	close(intchan)
}

// 从管道里面读取数据
func PrimeNum(intchan chan int, priChan chan int, exitChan chan bool) {
	defer wg1.Done()
	for num := range intchan {
		var falg = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				falg = false
				break
			}
		}
		if falg {
			priChan <- num // nums 是一个素数
		}
	}
	exitChan <- true
}

// 打印素数字
func PinrtFunction(priChan chan int) {
	defer wg1.Done()
	for v := range priChan {
		fmt.Printf("%v\n", v)
	}
}

func main() {
	star := time.Now().Unix()
	intChan := make(chan int, 1000)
	priChan := make(chan int, 1000)
	exitChan := make(chan bool, 16) // 标识priChan的数据

	wg1.Add(1)
	go PutNum(intChan)

	for i := 0; i < 16; i++ {
		wg1.Add(1)
		go PrimeNum(intChan, priChan, exitChan)
	}

	wg1.Add(1)
	go PinrtFunction(priChan)

	wg1.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(priChan)
		defer wg1.Done()
	}()
	wg1.Wait()
	end := time.Now().Unix()
	fmt.Println(end - star)
}
