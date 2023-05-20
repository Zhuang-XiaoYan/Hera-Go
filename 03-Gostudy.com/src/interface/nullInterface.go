package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("type :%T vayle %v\n", a, a)
}

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string) // 类型的断言 表示断言为string类型 但是不一定是正确的需要看ok的值 表示判断a的类型具体的什么操作 实际中可以更具子啊传入的类型 来实现对应操作。
	if !ok {
		fmt.Printf("猜错了")
	} else {
		fmt.Printf("猜对了 %v", str)
	}
}

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	// 使用x.(type) 来获取x 的类型是什么这样的 不是断言为什么类型 这样的写法只能是在switch 写法 不能在其他的地方使用
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

// 空接口 表示没有任何约束 任意的类型都是可以实现空接口

type A interface {
}

func main() {

	var m1 map[string]interface{}

	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 20
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	fmt.Println("------------------------------")
	show(false)
	show(nil)
	show(m1)
	assign(100)
	fmt.Println("----------------------------------------------------------")
	var a A
	var str = "你好"
	a = str
	fmt.Printf("value: %v tyep %T\n", a, a)

	var num = 20
	a = num
	fmt.Printf("value: %v tyep %T\n", a, a)

	var flag = true
	a = flag
	fmt.Printf("value: %v tyep %T\n", a, a)
	fmt.Println("----------------------------------------------------------")
	var mm2 = make(map[string]interface{})
	mm2["name"] = "zhangsan"
	mm2["age"] = 20
	fmt.Println(mm2)

	var ss1 = []interface{}{"hahah", 1, true}
	fmt.Println(ss1)
}
