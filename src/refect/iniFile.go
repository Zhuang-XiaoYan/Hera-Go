package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"address"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 判断的参数的必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("v param should be a pointer")
		return
	}
	// 传进来的v 必须是结构体类型的指针（因为配置的文件名中各种键值对需要复制给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("v param should be a struct")
		return
	}
	// 读文件得到字节类型
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 一行一行的读数据
	var structName string
	lineSlince := strings.Split(string(b), "\r\n")
	// 如果注释就跳过
	for index, line := range lineSlince {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果是[] 开头就是表示节
		if strings.HasPrefix(line, "[") {
			if line[0] == '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line :%d syntax error", index+1)
				return
			}
			// 把首位去掉[] 取到中间的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line :%d syntax error", index+1)
				return
			}
			// 根据字符串seactionName去data里面更具反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的结构体
					structName = field.Name
					fmt.Printf("找到了%s对应的嵌套结构体%s", sectionName, structName)
				}
			}
		} else {
			// 如果不是[] 就是表示=分割的键值对
			// 以等号分割这一行，等号左边是key 右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntanx error", index+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 更具structname 去data里面把对应的嵌套的结构体取出来
			v := reflect.ValueOf(data)
			structObj := v.Elem().FieldByName(structName)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s应该是一个结构体", structName)
				return
			}
			var filedName string
			var fileType reflect.StructField
			for i := 0; i < structObj.NumField(); i++ {
				filed := sType.Field(i) // tag信息是存储在类型信息中
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到了对应的字段
					filedName = filed.Name
					break
				}
			}
			if len(filedName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(filedName)
			fmt.Printf(filedName, fileType.Type)
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntanx error", index+1)
					return
				}
				fileObj.SetInt(valueInt)
			}
		}
	}
	return
}

func main() {
	var conf Config
	err := loadIni("./conf.ini", &conf)
	if err != nil {
		fmt.Printf("load ini failed ,error:%v\n", err)
	}
	fmt.Println(conf)
}
