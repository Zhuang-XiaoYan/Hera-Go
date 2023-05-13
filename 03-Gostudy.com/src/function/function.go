package main

import (
	"fmt"
	"sort"
)

// 函数就是一段代码的封装 这样可以重用

func sum(x int, y int) (ret int) {
	return x + y
}

func sum1(x int, y int) {
	fmt.Println(x + y + 10)
	fmt.Println("没有返回值")
}

// 没有参数的 但是有返回值  返回值的可与没有名
func sum2() int {
	return 0
}

func f3() int {
	res := 2
	return res
}

func f4(x int, y int) (res int) {
	res = x * y
	return // 使用的命名的返回值可以省略
}

func f5() (int, int) {
	return 10, 8
}

func f6(x, y, z int) (a, b int) {
	return x + y, y * z
}

// 当参数中的连续多个参数的类型的一致性的 我可以将非最后一个的参数类型的省略
func f7(x, y, z int, s1, s2 string) (a, b int) {
	return x + y, y * z
}

// 可边长的参数必须放在函数的最后面
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y类型的切片
}

func sumF1(x ...int) {
	fmt.Printf("%v--%T\n", x, x)
}

func sumF2(x, y int) {
	fmt.Printf("%v--%v\n", x, y)
}

// 可变参数
func sumF3(num ...int) (sum int) {
	for _v := range num {
		sum += _v
	}
	return sum
}

func mapSort(map1 map[string]string) string {
	var slicekey []string

	for key, _ := range map1 {
		slicekey = append(slicekey, key)
	}
	sort.Strings(slicekey)

	var str string
	for _, v := range slicekey {
		str += fmt.Sprint("%v=?%v", v, map1[v])
	}
	return str
}

// num ...int 只能放在最后的数据
func sumF4(x int, num ...int) int {
	sum := x
	for _v := range num {
		sum += _v
	}
	return sum
}

// 类型的升序排序
func sortIntASC(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}
	}
	return slice
}

// 定义在全局的便变量是是去哪句变量  在函数内部的变量是的局部变量。
var varsteing = "全局变量"

// 表示顶一个calc 的类型
type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

// 自定义一个方法类型
type calcType func(int, int) int

// 在函数中传递使用一个自定义类型参数
func calcNew(x, y int, cb calcType) int {
	return cb(x, y)
}

func sub(x, y int) int {
	return x - y
}

type cala func(int, int) int

func do(o string) cala {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		// 使用了匿名函数类型
		return func(x int, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func SortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}
	}
	return slice
}

func SortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] <= slice[j] {
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}
	}
	return slice
}

// 实现map的安装key的升序
func mapSortAsc(map1 map[string]string) string {
	var keySlice []string
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	sort.Strings(keySlice)
	var str string
	for _, v := range keySlice {
		str += fmt.Sprintf("key=%v, value=%v\n", v, map1[v])
	}
	return str
}

func fn1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	// 函数的递归调用
	fn1(n)
}

func main() {

	fn1(10)

	// 匿名函数自执行…………
	func() {
		fmt.Println("匿名函数自执行…………")
	}()

	var sliceA = []int{13, 15, 17, 19, 10}

	AscArr := SortIntAsc(sliceA)

	for _, v := range AscArr {
		fmt.Println(v)
	}

	fmt.Println("----------------------------------------")
	DesArr := SortIntDesc(sliceA)

	for _, v := range DesArr {
		fmt.Println(v)
	}

	fmt.Println("----------------------------------------")

	var c calc
	c = add
	fmt.Printf("c 的类型: %T\n", c)
	fmt.Printf("%d\n", c(10, 5))

	fmt.Println("-----------c 的类型-----------")
	var sliceC = []int{12, 34, 37, 35, 556, 36, 2}

	arr := sortIntASC(sliceC)

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("-----------函数参数的简写-----------")
	sumF1(12, 14, 56, 47, 28)
	fmt.Println("-----------函数参数的简写-----------")

	// 函数的定义
	fmt.Println(sum(1, 2))
	sum1(1, 4)
	fmt.Println(sum2())
	_, x := f5()
	fmt.Println(x)
	x1, x2 := f6(3, 5, 7)
	fmt.Println(x1, x2)

	f8("xjl", 6, 7)
	f8("xjl")
}
