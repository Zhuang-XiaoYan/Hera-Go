package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// os.O_WRONLY   只写权限
// os.O_RDONLY   只读权限
// os.O_RDWR     读写权限
// os.O_CREATE   创建文件
// os.O_APPEND   追加文件
// os.O_TRUNC    清空文件

func OsWriteFile() {
	file, err := os.OpenFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 00666)
	defer file.Close() // 及时关闭文件，防止内存泄露
	if err != nil {
		fmt.Println("file open error")
		return
	}
	var str = "SBSBSBSBSSBBSBSB\r\n"
	file.Write([]byte(str)) // 写入字节切片数据
	file.WriteString(str)   // 直接写入字符串数据
}
func BuferWriteFile() {
	file, err := os.OpenFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 00666)
	defer file.Close() // 及时关闭文件，防止内存泄露
	if err != nil {
		fmt.Println("file open error")
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("hell golang")
	// 将缓存中的内容写入到文件中
	writer.Flush()
}

func IoutilwriteFile() {
	str := "hello SB"
	err := ioutil.WriteFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("file open error")
		return
	}
}
func main() {
	OsWriteFile()
	BuferWriteFile()
	IoutilwriteFile()
}
