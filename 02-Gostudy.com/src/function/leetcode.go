package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 计算字符串中的汉字的数量
func countNumber(str string) int {
	var count = 0
	for _, value := range str {
		if unicode.Is(unicode.Han, value) {
			count++
		}
	}
	return count
}

// 单词出现的次数
func wordTime(str string) map[string]int {
	m1 := make(map[string]int, 10)
	str1 := strings.Split(str, " ")
	for _, vaule := range str1 {
		if _, ok := m1[vaule]; !ok {
			m1[vaule] = 1
		} else {
			m1[vaule]++
		}
	}
	return m1
}

func huiWen(str string) bool {
	// 先把字符串放在[]rune中
	r := make([]rune, 0, len(str))
	for _, c := range str {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "hihihadafha我是庄小焱"
	fmt.Println(countNumber(str))
	fmt.Println("-------------------------------------")
	var m = wordTime("how do you do")
	fmt.Printf("%T\n", m)
	for key, value := range m {
		fmt.Println(key, "==", value)
	}
	fmt.Println("-------------------------------------")
	var res = huiWen("山西运煤车车煤运西山")
	fmt.Println(res)
}
