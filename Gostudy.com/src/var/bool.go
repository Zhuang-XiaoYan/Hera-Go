package main

import (
	"fmt"
)

func main() {
	b1 := true
	fmt.Printf("%T\n", b1)
	var b2 bool // 默认是false
	fmt.Println(b2)
}
