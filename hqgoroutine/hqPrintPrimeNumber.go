package hqgoroutine

import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup

func HqPrintPrimeNumber()  {
	//runtime.GOMAXPROCS(1)

	//给每一个可用的内核分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)
	//创建两个goroutine
	go printPrime("A")
	go printPrime("B")

	//等待goroutine结束
	wg.Wait()
	fmt.Println("所有goroutine已结束")

}

func printPrime(prefix string)  {
	defer wg.Done()
	next:
		for outer := 2;outer<100 ; outer++{
			for inner := 2; inner<outer;inner++  {
				if outer % inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n",prefix,outer)
		}
		fmt.Println("打印完成",prefix)
}