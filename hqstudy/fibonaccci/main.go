package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"os"
)


type intGen func() int

//intGen实现Read接口进行使其可以当作文件来读
func (g intGen) Read(p []byte) (n int,err error)  {

	next := g()
	if next > 400{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	// TODO:incorrect if p is too small!
	//将strings.NewReader(s)临时存储一下，读完p后在返回
	return  strings.NewReader(s).Read(p)
}

//斐波那契数列
// 1 1 2 3 5 8 13 21
//   a b
//		a,b
func fibonacci() intGen {

	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}


func printfileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func writeFile(filename string)  {
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//defer writer.Flush()
	f := fibonacci()
	for i := 0; i <20 ;i++  {
		fmt.Fprintln(writer,f())
	}
}


func main() {

	f := fibonacci()
	fmt.Println(f())
	/*
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	*/
	//printfileContents(f)
	writeFile("fib.txt")


}
