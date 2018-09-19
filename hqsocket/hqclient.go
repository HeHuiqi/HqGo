package hqsocket

import (
	"net"
	"fmt"
)

func HqClient()  {

	ipStr := "127.0.0.1"
	addr := net.ParseIP(ipStr)
	if addr == nil {
		fmt.Println("ip 无效")
	}else {
		fmt.Println("IP为：",addr.String())
	}



}