package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	// a = b //不可以这样做，因为此时a和b是不同的类型
	fmt.Println(a, b)
	// 数组的定义 基本上与Java的定义是类似的定义
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
	fmt.Println("------------------------------------------------------------------------")
	// 按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string
}
