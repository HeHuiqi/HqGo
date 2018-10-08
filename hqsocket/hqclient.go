package hqsocket

import (
	"net"
	"fmt"
	"os"
)

func HqClient()  {

	/*
	if len(os.Args) !=2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		//os.Exit(1)
	}

	service := os.Args[1]
	*/

	service := "127.0.0.1:8080"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)

	if err != nil {
		fmt.Println("1-------")
		return
	}
	checkError(err)

	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		fmt.Println("2-------")
		return
	}
	checkError(err)
	//发送HTTP请求
	_,err = conn.Write([]byte("GET / HTTP/1.0 \r\n\r\n"))
	if err != nil {
		fmt.Println("3-------")
		return
	}
	checkError(err)

	rAddr := conn.RemoteAddr()
	//设置缓冲去读取数据
	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("4-------")
		return
	}
	checkError(err)
	fmt.Printf("Server:%s \nMessage:%s\n",rAddr.String(),string(buf[0:n]))
}
func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "HHQ:Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
