
package hqrpc

import (
"net/rpc"
"log"
"fmt"
"net/rpc/jsonrpc"
)



func HqRPCClient() {

	//HqHTTPRPCInvoke()
	//HqTCPRPCInvoke()
	HqJSONRPCInvoke()
}

func HqHTTPRPCInvoke()  {
	//服务端的地址和端口
	serverAddress := "127.0.0.1:9999"
	//创建一个RPC客户端
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}


	// Synchronous call
	args := Args{17, 8}

	var reply int
	//调用RPC服务
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith.Multiply error:", err) }
	fmt.Printf("Arith.Multiply: %d*%d=%d\n", args.A, args.B, reply)


	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("Arith.Divide error:", err) }
	fmt.Printf("Arith.Divide: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
func HqTCPRPCInvoke()  {
	serverAddress := "127.0.0.1:9999"
	//创建一个RPC客户端
	client, err := rpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}


	// Synchronous call
	args := Args{17, 8}

	var reply int
	//调用RPC服务
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith.Multiply error:", err) }
	fmt.Printf("Arith.Multiply: %d*%d=%d\n", args.A, args.B, reply)


	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("Arith.Divide error:", err) }
	fmt.Printf("Arith.Divide: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
func HqJSONRPCInvoke()  {
	serverAddress := "127.0.0.1:9999"
	//创建一个RPC客户端,注意区别TCP
	client, err := jsonrpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}


	// Synchronous call
	args := Args{17, 8}

	var reply int
	//调用RPC服务
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith.Multiply error:", err) }
	fmt.Printf("Arith.Multiply: %d*%d=%d\n", args.A, args.B, reply)


	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("Arith.Divide error:", err) }
	fmt.Printf("Arith.Divide: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}