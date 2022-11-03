package main

import "fmt"

// uint8类型，或者叫 byte 型，代表一个ASCII码字符。
// rune类型，代表一个 UTF-8字符。
func main() {

	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println("--------------------------------")
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	fmt.Println("--------------------------------")
	s2 := "白萝卜"
	s3 := []rune(s2)        // 字符串强制转换了一个rune的一个切片 ['白','萝','卜']
	s3[0] = '黑'             // 修改字符串
	fmt.Println(string(s3)) // 将的['黑','萝','卜'] 转为string "黑萝卜"
	fmt.Println("--------------------------------")
	// 强制类型转化
	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)

}
