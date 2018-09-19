package hqgoroutine

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
	"time"
)

var (
	counter int64

	hqWg sync.WaitGroup

	shutdown int64

	//mutex 用来定义一段代码临界区
	mutex sync.Mutex

)

func HqCounter()  {
	//计数器加2，表示要等待两个goroutine
	hqWg.Add(2)
	//创建两个goroutine
	go incCounter1(1)
	go incCounter1(2)

	//等待goroutine完成
	hqWg.Wait()
	fmt.Println("counter:",counter)

}
//原子操作保证同步
func incCounter1(id int)  {
	defer hqWg.Done()

	for count := 0;count<3 ;count++  {
		//原子操作使counter+1，如果不是原子操作，后一个goroutine会覆盖上一个
		atomic.AddInt64(&counter,1)
		//当前goroutine从线程退出，并放回队列
		runtime.Gosched()
	}
}


func HqCounter1()  {
	//计数器加2，表示要等待两个goroutine
	hqWg.Add(2)
	//创建两个goroutine
	go incCounter2(1)
	go incCounter2(2)

	//等待goroutine完成
	hqWg.Wait()
	fmt.Println("counter:",counter)

}
//使用互斥锁保证同步并安全访问
func incCounter2(id int)  {
	//在函数退出是调用通知main函数工作已经完成
	defer hqWg.Done()

	for count := 0;count<3 ;count++  {

		//同一时刻只允许一个goroutine进入临界区
		//临界区mutex.Lock()和mutex.Unlock()之间的代码就是临界区
		mutex.Lock()
		{
			value := counter

			//当前goroutine从线程退出，并放回到队列
			runtime.Gosched()

			value++
			counter = value
		}
		mutex.Unlock()

	}
}

/*
原子函数 LoadInt64()读取 和 StoreInt64()写入
这两个函数提供了一种安全地读 和写一个整型值的方式
*/
func HqDoWork()  {
	//计数器加2，表示要等待两个goroutine
	hqWg.Add(2)
	//创建两个goroutine
	go doWork("A")
	go doWork("B")

	//给定执行时间
	time.Sleep(1*time.Second)


	//该停止工作了
	fmt.Println("停止工作")
	atomic.StoreInt64(&shutdown,1)

}
func doWork(name string)  {

	defer  hqWg.Done()

	for   {
		fmt.Printf("Doing %s Work\n",name)
		//250毫秒
		time.Sleep(250*time.Millisecond)

		//要停止工作了吗
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutdown %s Down\n",name)
			break
		}
	}

}