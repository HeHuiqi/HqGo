package hqbase

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		fmt.Printf("%d", v)
		total += v
	}
	c <- total
}

func HqArrayUse()  {
	fmt.Println("hello")
	//创建并初始化一个数组
	a := []int{1, 2, 3, 4, 5, 6}
	b := [...]int{4,5,6,7}
	fmt.Println(b)

	//创建一个无缓冲的通道
	c := make(chan int)
	//开启两个goroutine
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	//读取各个goroutine在通道存储的值
	x, y := <-c, <-c
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("x+y=", x+y)

}
func HqSliceUse()  {

	//创建一个int类型slice切片
	hqBooks := make([]int,3,5)
	hqBooks[0] = 9
	hqBooks[1] = 10
	fmt.Println("hqBooks1=",hqBooks)
	//同过切出一个子切片来间接删除元素
	index := 0
	hqBooks =  append(hqBooks[:index],hqBooks[index+1:]...)
	fmt.Println("hqBooks2=",hqBooks)

	//通过字面量创建一个string类型的slice
	teachers := []string{"李老师","刘老师","夏老师","陆老师"}
	teachers = append(teachers,"张老师")
	teachers = append(teachers,"徐老师")
	for _,value := range teachers {
		fmt.Println(value)
	}
}