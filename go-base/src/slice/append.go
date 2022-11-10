package main

import "fmt"

func main() {
	// append 是为切片追加元素
	s1 := []string{"上海", "北京", "深圳"}
	// s1[3] = "广州" // 数组的越界
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// 调用append函数必须使用原来的切片的变量来接受返回值
	s1 = append(s1, "江西")
	fmt.Println(s1) // append追加元素的时候 原来的底层数组不够的时候 go-base 语言的底层会将底层数组也会换一个底层数组
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "成都", "重庆")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := []string{"武汉", "长沙", "苏州"}
	s1 = append(s1, s2...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，
	// 此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。

	// newcap := old.cap
	// doublecap := newcap + newcap
	// if cap > doublecap {
	//	newcap = cap
	// } else {
	//	if old.len < 1024 {
	//		newcap = doublecap
	//	} else {
	//		// Check 0 < newcap to detect overflow
	//		// and prevent an infinite loop.
	//		for 0 < newcap && newcap < cap {
	//			newcap += newcap / 4
	//		}
	//		// Set newcap to the requested cap when
	//		// the newcap calculation overflowed.
	//		if newcap <= 0 {
	//			newcap = cap
	//		}
	//	}
	//}

	// 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	// 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	// 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	// 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	fmt.Println("-----------------------------------------------------")
	// 切片拷贝  由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
	a := []int{1, 2, 3, 4, 5}
	b := a         // b a 同时指向同一个的内存地址
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
