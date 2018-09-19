package hqbase

import "fmt"

//定义接口
type HqBehaviour interface {
	SayHi()
	Sing(lyric string)
}
type Person struct {
	name string
	age  int
}

func (person Person)EatFood()  {
	fmt.Println(person.name,"正在吃东西。。。")
}

//实现接口SayHi方法
func (person Person) SayHi() {

	fmt.Println(person.name, "你好")

}

//实现接口Sing方法
func (person Person) Sing(lyric string) {
	fmt.Println(person.name, "正在唱歌。。。", lyric)
}

//Student包含了一个Person匿名变量，他就拥有了Person所有的属性和方法包括接口实现
type Student struct {
	Person //使用一个匿名变量即可实现继承的作用
	idNo   string
	class  string
}

func (student Student) DoingHomework() {
	fmt.Println(student.name, "正在做家庭作业......")
}


func HqOOPUse()  {

	//字面量创建一个对象
	person :=  Person{"小刘", 20}
	person.SayHi()
	person.Sing("十八弯的水路到我家")

	//字面量创建一个对象,写在一行需要初始化所有参数
	student := Student{Person{"韩梅梅", 17}, "95270", "一年级一班"}
	student.SayHi()
	student.Sing("让我们荡起双桨")
	student.DoingHomework()


	//结构化创建对象,初始化参数可选，需要填写属性（参数）名称
	bigStu := Student{
		idNo:"12315",
	}
	bigStu.name = "明景泰"
	bigStu.EatFood()//调用了Person的方法
	fmt.Println(bigStu)


	var hqBehaviour HqBehaviour
	//接口类型指向实现自己方法的对象（实例）
	hqBehaviour = Person{"呵呵呵", 1}
	hqBehaviour.SayHi()

	//接口类型指向实现自己方法的对象（实例）
	hqBehaviour = Student{Person{"哈哈哈", 10}, "3333", "三班"}
	hqBehaviour.Sing("kkkh")


}
