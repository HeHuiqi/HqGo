package main

import (
	"fmt"
	"testing"
)

/*

Go 语言的测试工具只会认为以_test.go 结尾的文件是测试文件
如果没有遵从这个约定，在包里 运行go test的时候就可能会报告没有测试文件。
一旦测试工具找到了测试文件，就会查找里 面的测试函数并执行

一个测试函数 必须是公开的函数，并且以 Test 单词开头。不但函数名字要以 Test 开头，
而且函数的签名必须接收一个指向 testing.T 类型的指针，并且不返回任何值。
如果没有遵守这些约定，测试框 架就不会认为这个函数是一个测试函数，也不会让测试工具去执行它
如：func TestXXX(t *testing.T) { ... }
执行命令
go test -v
*/
func TestHello(t *testing.T) {
	fmt.Println("hello")

}
func HqRequest()  {
	for i := 0; i < 100; i++ {
		//fmt.Println(i)
	}
}
/*
已Benchmark开头的函数是做性能(压力)测试的b.N表示最大测试次数
如：func BenchmarkXXX(b *testing.B) { ... }

go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，
语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
我们执行命令go test -bench=".*" -count=5（使用-count可以指定执行多少次）
执行命令
go test -bench=.


*/
func BenchmarkRequest(b *testing.B)  {
	//fmt.Println("b.N=",b.N)
	for n := 0;n<b.N ; n++ {
		HqRequest()
	}
}