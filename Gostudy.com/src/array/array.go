package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

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
	// 按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，... 可以推断来是实现 与的java中的Arraylist是一样的意思
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string
	fmt.Println("------------------------------------------------------------------------")
	// 我们还可以使用指定索引值的方式来初始化数组
	a2 := [...]int{1: 1, 3: 5}
	fmt.Println(a2)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a2) //type of a:[4]int
	fmt.Println("------------------------------------------------------------------------")
	//数组的遍历
	var a3 = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	// 方法2：for range遍历
	for index, value := range a3 {
		fmt.Println(index, value)
	}
	// 二维数组的定义
	a4 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a4)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a4[2][1]) //支持索引取值:重庆

	// 二维数组的遍历  for range 与java中的foreach相同的工作
	for _, v1 := range a4 {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Println()
	}

	//  多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
	//支持的写法
	a5 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	b5 := [3][...]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a5, b5)

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。 这个与java是不同的，在Java中是引用传递，而go中是值传递
	a6 := [3]int{10, 20, 30}
	modifyArray(a6) //在modify中修改的是a的副本x
	fmt.Println(a6) //[10 20 30]
	b6 := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b6) //在modify中修改的是b的副本x
	fmt.Println(b6)  //[[1 1] [1 1] [1 1]]
}
