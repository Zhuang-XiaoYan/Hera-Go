package main

import (
	"fmt"
	"strconv"
)

// strconv
// 把字符串解析成为的数字类型 有点像的java中的Integer.parseint() 函数

func main() {
	str := "10000"
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Errorf("error")
	}
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("%T %v", res, res)
	// 直接将字符串转为int类型的数字
	res3, _ := strconv.Atoi("600")
	fmt.Println()
	fmt.Printf("%T %d", res3, res3)
	fmt.Println()
	fmt.Printf("--------------------------------------------------\n")
	// 把数字转为的字符串类型
	i := int32(97)
	res2 := fmt.Sprintf("%d", i)
	fmt.Printf("%T %v", res2, res2)
	fmt.Printf("--------------------------------------------------\n")
	boolres := "true"
	boolvalue, _ := strconv.ParseBool(boolres)
	fmt.Printf("%T %v", boolvalue, boolvalue)
	fmt.Printf("--------------------------------------------------\n")
	floatres := "3.1415926"
	floatvalue, _ := strconv.ParseFloat(floatres, 64)
	fmt.Printf("%T %v", floatvalue, floatvalue)
}
