package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("type :%T vayle %v\n", a, a)
}

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Printf("猜错了")
	} else {
		fmt.Printf("猜对了 %v", str)
	}
}

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
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

}
