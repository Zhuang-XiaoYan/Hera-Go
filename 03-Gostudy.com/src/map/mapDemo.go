package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 使用for range 遍历的都是key value的值 如果是map 类型那就是访问key value的类型  如果是数组  那就是index value 的值。

func cal() {
	var str = "how do you do"
	splits := strings.Split(str, " ")
	fmt.Println(splits)
	fmt.Println("-----------------------------")
	var strMap = make(map[string]int)
	for _, v := range splits {
		strMap[v] += 1
	}
	fmt.Println(strMap)
}

func main() {
	cal()

	// map 相当于java中的hashmap
	var hash map[string]int
	hash = make(map[string]int) // 初始化 map的申明 相当于java中HashMap map=new HashMap()
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

	// 这个是map的声明并初始化的定义
	var userinfo = map[string]string{
		"username": "zhangxiaoyan",
		"age":      "20",
		"sex":      "男",
	}

	// 使用的是引用推导的功能来定义
	userdata := map[string]string{
		"username": "zhangxiaoyan",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println("--------------------------------------")
	fmt.Printf("map的初始化定义为：%v\n", userdata)
	fmt.Printf("map的初始化定义为:%v\n", userinfo)
	fmt.Println("--------------------------------------")
	// map的遍历
	for key, value := range userdata {
		fmt.Println(key, "==", value)
	}

	v, ok := userdata["test"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在 这个属性")
	}
	fmt.Println("--------------------------------------")
	// 删除
	fmt.Println(hash)
	// 删除map中一个key-value值
	delete(hash, "xjl")
	fmt.Println(hash)

	fmt.Println("-------------------------这是使用slince的map类型---------------------")

	var user1 = make([]map[string]string, 2, 3)

	if user1[0] == nil {
		user1[0] = make(map[string]string)
		user1[0]["name"] = "xiaoyan"
		user1[0]["age"] = "27"
		user1[0]["sex"] = "man"
	}
	if user1[1] == nil {
		user1[1] = make(map[string]string)
		user1[1]["name"] = "丫头"
		user1[1]["age"] = "25"
		user1[1]["sex"] = "woman"
	}

	for _, mapobj := range user1 {
		// 遍历每一个map的对象
		//for key, value := range v {
		//	// 便利每一个map的属性
		//	fmt.Printf("key= %v==value%v\n", key, value)
		//}
		fmt.Printf("%v\n", mapobj)
	}
	fmt.Println("-------------------------这是使用slince的map类型---------------------")

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
	var userinfo1 = make(map[string][]string)

	userinfo1["habby"] = []string{
		"吃饭",
		"睡觉",
		"拉屎",
	}

	userinfo1["work"] = []string{
		"java",
		"python",
		"golang",
	}

	// golang 中使用range 来便利map  就是key value的像是 那如果 for i 来实现 那就是用过下标来实现。
	for k, v := range userinfo1 {
		fmt.Printf("%v ,%v\n", k, v)
	}

	//fmt.Println(userinfo1)
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

	fmt.Println("---------------------值类型与引用类型-----------------------------------------------")
	// 声明的是一个map 类型变量 key value 都是string 类型
	var userdata3 = make(map[string]string)

	userdata3["username"] = "张三"
	userdata3["age"] = "20"

	userdata4 := userdata3

	fmt.Println(userdata3)
	fmt.Println(userdata4)

	userdata4["username"] = "xiaoyan"
	// map类型是一个引用类型变量
	fmt.Println(userdata3)
	fmt.Println(userdata4)

	fmt.Println("---------------------map key 升序排列-----------------------------------------------")
	map2 := make(map[int]int, 10)
	map2[10] = 9
	map2[1] = 77
	map2[5] = 6
	map2[4] = 90
	map2[3] = 66
	// 这里是随意的按照key的顺序来实现打印
	for key, val := range map2 {
		fmt.Println(key, val)
	}

	// 使用升序来是新
	// 把map中的key 放在切片里面
	var keySlice []int
	for key, _ := range map2 {
		keySlice = append(keySlice, key)
	}
	sort.Ints(keySlice)
	fmt.Println(keySlice)

	for _, v := range keySlice {
		fmt.Printf("key=%v, value=%v\n", v, map2[v])
	}
}
