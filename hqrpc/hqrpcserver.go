package hqrpc

import (
	"errors"
	"net/rpc"
	"net/http"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type Arith int

/*

RPC 就是远程过程调用，通过http或tcp协议来调用服务的方法，如这里的乘法和除法以操作方法

Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下:
函数必须是导出的(首字母大写)
必须有两个导出类型的参数，
第一个参数是接收的参数，
第二个参数是返回给客户端的参数，必须是指针类型的
函数还要有一个返回值error

举个例子，正确的RPC函数格式如下:
func (t *T) MethodName(argType T1, replyType *T2) error
T、T1和T2类型必须能被encoding/gob包编解码。

*/
func (t *Arith) Multiply (args *Args ,reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args , quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = args.A /args.B
	quo.Rem = args.A % args.B
	return nil
}
func HqRPCServer()  {

	//hqHTTPRPC()
	//HqTCPRPC()
	hqJSONRPC()
}

func hqHTTPRPC()  {
	arith := new(Arith)
	rpc.Register(arith)
	//建立在http协议的RPC服务
	rpc.HandleHTTP()
	fmt.Println("httprpc","127.0.0.1:9999")

	err := http.ListenAndServe("127.0.0.1:9999",nil)
	if err != nil {

		fmt.Println("HH:",err.Error())
	}
}
func HqTCPRPC()  {
	arith := new(Arith)
	//注册prc服务
	rpc.Register(arith)

	tcpAddr ,err := net.ResolveTCPAddr("tcp","127.0.0.1:9999")
	if err !=nil {
		fmt.Println("1---",err)
		return
	}
	fmt.Println("tcprpc","127.0.0.1:9999")

	//建立在TCP协议的RPC服务
	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err !=nil {
		fmt.Println("2---",err)
		return
	}
	for  {
		conn,err := listener.Accept()

		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
//JSON RPC是数据编码采用了JSON，而不是gob编码
func hqJSONRPC()  {
	arith := new(Arith)
	//注册prc服务
	rpc.Register(arith)

	tcpAddr ,err := net.ResolveTCPAddr("tcp","127.0.0.1:9999")
	if err !=nil {
		fmt.Println("1---",err)
		return
	}
	fmt.Println("jsonrpc","127.0.0.1:9999")

	//建立在TCP协议的RPC服务
	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err !=nil {
		fmt.Println("2---",err)
		return
	}
	for  {
		conn,err := listener.Accept()

		if err != nil {
			continue
		}
		//注意这里和TCP协议的PRC的区别
		jsonrpc.ServeConn(conn)
	}
}