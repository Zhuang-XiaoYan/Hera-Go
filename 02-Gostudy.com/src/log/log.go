package main

import "log"

// 使用的自带的log库 logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。

//const (
//	// 控制输出日志信息的细节，不能控制输出的顺序和格式。
//	// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go-base:23: message
//	Ldate         = 1 << iota     // 日期：2009/01/23
//	Ltime                         // 时间：01:23:23
//	Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
//	Llongfile                     // 文件全路径名+行号： /a/b/c/d.go-base:23
//	Lshortfile                    // 文件名+行号：d.go-base:23（会覆盖掉Llongfile）
//	LUTC                          // 使用UTC时间
//	LstdFlags     = Ldate | Ltime // 标准logger的初始值
//)

func main() {
	//log.Println("这是一条很普通的日志。")
	//v := "很普通的"
	//log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")
	//log.Panicln("这是一条会触发panic的日志。")

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
