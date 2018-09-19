package hqstdlib

import (
	"fmt"
	"time"
)

func HqTimerUser()  {
	f := func() {
		fmt.Println("Time out")
	}
	time.AfterFunc(1*time.Second, f)
	time.Sleep(2 * time.Second)//要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
	//【结果】运行了1秒后，打印出timeout，又过了1秒，程序退出
	//将一个间隔和一个函数给AfterFunc后
	//间隔时间过后，执行传入的函数

	t1 := time.Now()

	t2 := t1.Add(time.Hour)
	fmt.Println("t2=",t2)

}