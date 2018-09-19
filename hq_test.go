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
*/
func TestHello(t *testing.T) {
	fmt.Println("hello")

}
