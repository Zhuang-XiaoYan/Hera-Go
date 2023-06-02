package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CpoyByIoutil() {
	file, err := ioutil.ReadFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test.txt")
	if err != nil {
		fmt.Println("file open err", err)
		return
	}
	err = ioutil.WriteFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\testcpoy.txt", file, 0666)
	if err != nil {
		fmt.Println("file wite err", err)
		return
	}
	fmt.Println(" file cpoy success")
}

func CpoyFile() {

	readfile, err := os.Open("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test.txt")
	if err != nil {
		fmt.Println("read file failed")
		return
	}
	openfile, err := os.OpenFile("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\test2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 00666)
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer readfile.Close() // 及时关闭流 防止内存泄露
	defer openfile.Close()

	var tmpFile = make([]byte, 128)
	for {
		n, err := readfile.Read(tmpFile)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = openfile.Write(tmpFile[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func MkDir() {
	// 也是可以使用os.Mkdir(filepath,perm) 如果存在会throw err
	err := os.MkdirAll("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\1", 00666) // 创建多级目录
	if err != nil {
		fmt.Println("create　dir failed", err)
	}
}

func RemoveDir() {
	// 也是可以使用os.Remove(filepath) 删除可以是删除文件或者目录 如果不存在会throw err
	err := os.RemoveAll("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\1") // 删除dir
	if err != nil {
		fmt.Println("remove　dir failed", err)
	}
}

func ReNameFile() {
	err := os.Rename("D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\1", "D:\\gitee\\Hera-Go\\03-Gostudy.com\\src\\file\\2")
	if err != nil {
		fmt.Println("Rename　dir failed", err)
	}
}

func main() {
	// CpoyByIoutil()
	//CpoyFile()
	//MkDir()
	ReNameFile()
}
