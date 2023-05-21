package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)
import "github.com/shopspring/decimal"

//Usage:   //go mod <command> [arguments]
//The commands are:
//    download    download modules to local cache
//    edit        edit go.mod from tools or scripts
//    graph       print module requirement graph
//    init        initialize new module in current directory
//    tidy        add missing and remove unused modules
//    vendor      make vendored copy of dependencies
//    verify      verify dependencies have expected content
//    why         explain why packages or modules are needed

const json = `{"name":{"first":"Janet","last":"zhangxiaoyan"},"age":47}`

func main() {

	var num float64 = 1
	for num < 10 {
		rs := num/100 + 1
		decimalRs := decimal.NewFromFloat(rs) //类似于java 中Bigdecimal 处理浮点数的运算方法
		fmt.Println(rs, decimalRs)
		fmt.Printf("origin:%v,originType:%T,decimal:%v,decimalType:%T\n", rs, rs, decimalRs, decimalRs)
		num += 1
	}

	fmt.Println("----------------------------------------------------")

	value := gjson.Get(json, "name.last")
	println(value.String())
}
