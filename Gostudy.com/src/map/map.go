package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// map 相当于java中的hashmap
	var hash map[string]int
	hash = make(map[string]int) // 初始化
	hash["xjl"] = 20
	hash["庄小焱"] = 30
	fmt.Println(hash) // 还么有在在内存开辟空间 导致没有初始化
	// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。

	value, ok := hash["xjl1"]
	if !ok {
		fmt.Println("没有这值")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for key, value := range hash {
		fmt.Println(key, "==", value)
	}
	// 删除
	fmt.Println(hash)
	delete(hash, "xjl")
	fmt.Println(hash)

	fmt.Println("---------------------------------------------------------------------")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(10)           //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	fmt.Println("---------------------------------------------------------------------")
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "xjl"
	fmt.Println(s1)

	fmt.Println("---------------------元素为map类型的切片------------------------------------------------")
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("---------------------值为切片类型的map------------------------------------------------")
	var sliceMap1 = make(map[string][]string, 3)
	fmt.Println(sliceMap1)
	fmt.Println("after init")
	key := "中国"
	value1, ok1 := sliceMap1[key]
	if !ok1 {
		value1 = make([]string, 0, 2)
	}
	value1 = append(value1, "北京", "上海")
	sliceMap1[key] = value1
	fmt.Println(sliceMap1)

}
