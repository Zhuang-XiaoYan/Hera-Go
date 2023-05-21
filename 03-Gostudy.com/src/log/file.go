package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件

// 使用for循环读取文件中的所有数据。
func read(path string) {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio是在file的基础上封装了一层API，支持更多的功能。
func readFileByBufferIo(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
func readFileByIoutil(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

// 使用最基础的code 来实现文件的读取
func baseReadFile(path string) {
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file failed error%v", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	for {
		// var tmp =[128]byte
		// fileObj.Read(tmp[:])
		n, error := fileObj.Read(tmp)
		if error != nil {
			fmt.Printf("read from file failed err%v", error)
			return
		}
		fmt.Printf("读了%d 字节", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
// os.O_WRONLY	只写
// os.O_CREATE	创建文件
// os.O_RDONLY	只读
// os.O_RDWR	读写
// os.O_TRUNC	清空
// os.O_APPEND	追加
func writeFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据
}

// bufio.NewWriter 来实现文件的写入功能
func writeBuffIo() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

// ioutil.WriteFile来实现的文件的写入
func writeIoUtils() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func main() {
	// var path = "D:\\gitee\\Hera-Go\\src\\log\\logDemo.go-base"
	// baseReadFile(path)
	//fmt.Println("-----------------------11-----------------------------")
	//readFileByIoutil(path)
	//fmt.Println("------------------------22----------------------------")
	//readFileByBufferIo(path)
	//fmt.Println("-------------------------33---------------------------")
	//read(path)
	//fmt.Println("-------------------------44---------------------------")

	var path2 = "D:\\gitee\\Hera-Go\\src\\log\\test.txt"
	writeFile(path2)
	fmt.Println("-----------------------copyFile--------------------------")

	_, err := CopyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}
