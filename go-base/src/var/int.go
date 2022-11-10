package main

import "fmt"

func main() {
	// 而且还允许我们用 _ 来分隔数字
	v := 123_456
	fmt.Println(v)

	i0 := 101     // 十进制
	i1 := 0777    // 八进制
	i2 := 0x13456 //十六类型
	// 十进制输出
	fmt.Printf("%d\n", i0)
	// 十进制转二进制
	fmt.Printf("%b\n")
	// 十进制转八进制
	fmt.Printf("%o\n", i0)
	// 十进制转十六进制
	fmt.Printf("%x\n", i1)
	fmt.Printf("%d\n", i2)
	// 插卡类型
	fmt.Printf("%T\n", i0)
	// 声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%d\n", i4)

	fmt.Printf("-------------------------------------------------")

}
