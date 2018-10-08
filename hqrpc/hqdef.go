package hqrpc

type Args struct {
	A, B int
}
//客户端的数据结构要和服务端的数据结构保持一致
type Quotient struct {
	Quo, Rem int
}