package main

import (
	"fmt"
	"time"
)

// timeDemo 时间对象的年月日时分秒
func timeDemo() {
	timeobj := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", timeobj)

	year := timeobj.Year()     // 年
	month := timeobj.Month()   // 月
	day := timeobj.Day()       // 日
	hour := timeobj.Hour()     // 小时
	minute := timeobj.Minute() // 分钟
	second := timeobj.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}
func timezoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

// timestampDemo 时间戳
func timestampDemo() {
	//now := time.Now()        // 获取当前时间
	//timestamp := now.Unix()  // 秒级时间戳
	//milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	//micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	//nano := now.UnixNano()   // 纳秒时间戳
	//fmt.Println(timestamp, milli, micro, nano)
}

//func timestamp2Time() {
//	// 获取北京时间所在的东八区时区对象
//	secondsEastOfUTC := int((8 * time.Hour).Seconds())
//	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
//
//	// 北京时间 2022-02-22 22:22:22.000000022 +0800 CST
//	t := time.Date(2022, 02, 22, 22, 22, 22, 22, beijing)
//
//	var (
//		sec  = t.Unix()
//		msec = t.UnixMilli()  // 毫秒时间戳 Go1.17+
//		usec = t.UnixMicro()  // 毫秒时间戳 Go1.17+
//	)
//
//	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
//	timeObj := time.Unix(sec, 22)
//	fmt.Println(timeObj)           // 2022-02-22 22:22:22.000000022 +0800 CST
//	timeObj = time.UnixMilli(msec) // 毫秒级时间戳转为时间对象
//	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
//	timeObj = time.UnixMicro(usec) // 微秒级时间戳转为时间对象
//	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
//}

// 时间的操作
// func (t Time) Add(d Duration) Time
// func (t Time) Sub(u Time) Duration
// func (t Time) Equal(u Time) bool
// func (t Time) Before(u Time) bool
// func (t Time) After(u Time) bool

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

// formatDemo 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}

func main() {
	timeDemo()

	timezoneDemo()
}
