package hqsocket

import (
	"net"
	"fmt"
)

func HqServer()  {

	service := ":8080"

	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	checkError(err)
	//监听端口
	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for  {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}

		//并发处理客户端请求
		go handleClient(conn)
	}
}

//func checkError(err error)  {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client：", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client!"))
		if err2 != nil {
			return
		}
	}
}