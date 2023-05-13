package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
	fmt.Println("-----------------------------------------")
	// copy()复制切片
	a1 := []int{1, 2, 3, 4, 5}
	c1 := make([]int, 5, 5)
	copy(c1, a1)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a1) //[1 2 3 4 5]
	fmt.Println(c1) //[1 2 3 4 5]
	c1[0] = 1000
	fmt.Println(a1) //[1 2 3 4 5]
	fmt.Println(c1) //[1000 2 3 4 5]
	fmt.Println("-----------------------------------------")
	// 从切片中删除元素
	a2 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Println(a2)
	// 要删除索引为2的元素
	a2 = append(a2[:2], a2[3:]...)
	fmt.Println(a2) //[30 31 33 34 35 36 37]
	fmt.Println("-----------------------------------------")
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1[:1], s1[2:]...) // 修改了底层数组
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1)
	fmt.Println(x1, len(x1), cap(x1))

	var a3 = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a3 = append(a3, i)
	}
	fmt.Println(a3, len(a3), cap(a3))

	var a4 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a4 = append(a4, fmt.Sprintf("%v", i))
	}
	fmt.Println(a4, len(a4), cap(a4))
	// 排序的函数实现
	sort.Ints(a3[:])
	//sort.Ints(a3)
	fmt.Println(a3)

	// append 的删除操作
	ss := [...]int{1, 3, 4, 5, 7, 9}
	ss1 := ss[:]
	ss1 = append(ss1[0:1], ss1[2:]...)
	fmt.Println(ss1)
	fmt.Println(ss)
}
