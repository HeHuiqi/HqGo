package hqstdlib

import (
	"bytes"
	"fmt"
	"os"
	"net/http"
	"log"
	"io"
)

func HqIOUse()  {


	//判断一个数据是否是某个类型，但是这个数据必须是一个接口声明的变量
	var sum interface{}
	sum = 10
	if value,ok := sum.(int);ok {
		fmt.Println(value)
	}

	fmt.Println("-----------------")

	//创建一个buffer
	var buf bytes.Buffer
	//向buffer中写入字符串
	buf.Write([]byte("Hello "))
	//Fprintf()会将一个字符串拼接到buffer
	fmt.Fprintf(&buf,"World!\n")
	//将buffer中的内容输出到标准输出设备
	buf.WriteTo(os.Stdout)
}
func HqCurl()  {

	r,err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file ,err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dest := io.MultiWriter(os.Stdout,file)

	io.Copy(dest,r.Body)
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}

