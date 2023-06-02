package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 如果是大量的文件 采用上面两种方法来读取文件
func OpenFile() {
	// 使用os下的open方法是只读的方法打开目录下的main.go文件
	file, err := os.Open("D:\\gitee\\Hera-Go\\02-Go语言开发环境构建\\Go语言开发环境构建.md")
	defer file.Close() // 及时关闭流 防止内存泄露
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	// 操作文件
	fmt.Println(file) // 打印的是文件的对象
	var strSlice []byte
	var tmpfile = make([]byte, 1280)
	for {
		n, err := file.Read(tmpfile)
		if err == io.EOF { // 表示读取完毕
			fmt.Println("read file finish")
			break
		}
		if err != nil {
			fmt.Println("read file failed")
			return
		}
		fmt.Printf("read %v byte", n)
		strSlice = append(strSlice, tmpfile[:n]...)
	}
	fmt.Println(string(strSlice)) // 只能读取部分文件
}

func BuferIoFile() {
	file, err := os.Open("D:\\gitee\\Hera-Go\\02-Go语言开发环境构建\\Go语言开发环境构建.md")
	defer file.Close() // 及时关闭流 防止内存泄露
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	reader := bufio.NewReader(file)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read file failed")
		return
	}
	var fileContext string
	// 读取一行数据
	fmt.Println(readString)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF { // 表示读取完毕
			fileContext += readString
			fmt.Println("read file finish")
			break
		}
		fileContext += readString
	}
	fmt.Println(fileContext)
}

// 在小文件的时候可以采用的方法
func IoUtilFile() {
	file, err := ioutil.ReadFile("D:\\gitee\\Hera-Go\\02-Go语言开发环境构建\\Go语言开发环境构建.md")
	if err != nil {
		fmt.Println("read file failed ", err)
		return
	}
	fmt.Println(string(file))
}

func main() {
	//OpenFile()
	fmt.Println("-----------------------------------")
	// BuferIoFile()
	fmt.Println("-----------------------------------")
	IoUtilFile()
}
