package hqgoroutine

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"strconv"
)

//创建通道需要使用make()函数
/*
无缓冲的整型通道
unbuffer := make(chan int)
//有缓冲的字符串通道,第二个参数表示缓冲区的大小
bufferstr := make(chan string,10)
 */
var hqChanWg  sync.WaitGroup
func init()  {
	rand.Seed(time.Now().UnixNano())
}
//无缓冲的通道指在接收前没有能力保存任何值的通道
//这种类型的通 道要求发送 goroutine 和接收 goroutine 同时准备好，
// 才能完成发送和接收操作。
func HqUnBufferChanel()  {

	 courtChan := make(chan int)
	 //计数器加2，表示要等待两个goroutine
	 hqChanWg.Add(2)

	 //启动两个选手
	 go player("AAA",courtChan)
	 go player("BBB",courtChan)

	 //发球
	 courtChan <- 1

	 //等待游戏结束
	 hqChanWg.Wait()
}
//模拟两个选手在打网球
func player(name string,counter chan  int)  {
	defer hqChanWg.Done()
	for{
		//等待通道被关闭，我们就赢了
		ball,ok := <-counter
		if !ok {
			fmt.Printf("Player %s Won\n",name)
			return
		}


		//选随机数，然后用这个数判断我们是否丢球
		n := rand.Intn(1000)
		if n%13 == 0 {
			fmt.Printf("---Player %s rand %d\n",name,n)
			fmt.Printf("Player %s Missed\n",name)
			//通道关闭表示我们输了
			close(counter)
			return
		}

		fmt.Printf("Player %s Hit %d\n",name,ball)
		//显示击球数，并将击球加1
		ball++
		//将球打向对手
		counter <- ball
	}
}


 //有缓冲的通道
func HqBufferChanel()  {

	money := make(chan int,10)
	//向通道写入一个值
	money <- 56

	value := <- money
	fmt.Println("通道的值为：",value)

}

func HqChanelUse()  {

	ch := make(chan int,1)
	timeout := make(chan bool,1)
	go func() {
		time.Sleep(time.Second*1)
		fmt.Println("超时了。。。。。。。。")
		timeout <- true
	}()

	str := "SS:"
	for i := 0;i<10;i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		case <-timeout:
			
		}
		j := <- ch
		str = str+strconv.Itoa(j)
	}
	fmt.Printf("%s\n",str)


	//这里模拟设置2秒才能保证超时调用
	time.Sleep(time.Second*2)
}
func HqGoPrint()  {

	t := time.Now()
	for i := 0; i < 10; i++ {
		 hqPrint(i)
	}
	fmt.Println("HqGoPrint--",t.Sub(time.Now()))


	//time.Sleep(time.Second*1)
}
func HqGoPrint1()  {
	chs := make([]chan int,10)
	t := time.Now()

	for i := 0; i < 10; i++ {
		chs[i] = make(chan  int)
		//go hqPrint(i)
		go hqPrintf(chs[i],i)
	}
	fmt.Println("HqGoPrint1--",t.Sub(time.Now()))
	for _,ch := range chs{
		<- ch
	}


	//time.Sleep(time.Second*1)
}
func hqPrintf(ch chan  int,index int)  {

	fmt.Printf("%d Hello World!F\n",index)
	ch <- 1
}
func hqPrint(index int)  {
	fmt.Println(index,"Hello World!")
}

/*
使用channel共享数据
*/
