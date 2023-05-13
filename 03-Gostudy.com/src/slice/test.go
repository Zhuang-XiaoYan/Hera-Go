package main

import (
	"fmt"
	"sort"
)

func main() {
	// 本文件就是实现的对排序算法的go 实现。
	var arr = [...]int{1, 5, 7, 3, 8}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("%v，%v \n", i, j)
			}
		}
	}
	fmt.Println("--------------------------------------------")
	var numSlice = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] <= numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)

	fmt.Println("---------------sort-----------------------------")
	intList := []int{2, 5, 7, 2, 3, 5, 7, 9}
	flostList := []float64{4.2, 4, 7, 9.8}
	stringList := []string{"wishi", "zhuang xiaoyan"}

	// 升序的排序
	sort.Ints(intList)
	sort.Float64s(flostList)
	sort.Strings(stringList)
	fmt.Printf("sort 排序后的结果是\n %v \n %v \n %v\n", intList, flostList, stringList)

	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(flostList)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

}
