package hqgoroutine

import (
	"runtime"
	"sync"
	"fmt"
)

func HqGoroutinePrintChar()  {

	//分配一个逻辑处理器给调度器使用功
	runtime.GOMAXPROCS(1)
	//wg表示等待程序完成，计数器加2表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		//显示字符3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26;char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()
	//fmt.Println("-------")
	//声明一个匿名函数，并创建一个goroutine
	go func() {
		defer  wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A';char < 'A'+26 ;char++  {
				fmt.Printf("%c ",char)

			}
		}
	}()

	//等待goroutine结束
	fmt.Println("等待任务完成")
	wg.Wait()
	fmt.Println("程序结束")
}
