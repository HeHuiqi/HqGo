package hqbase

import "fmt"

func HqMapUse()  {
	//同过make()函数创建一个key为int类型，value为string类型的map
	myMap := make(map[int]string)
	myMap[1] = "请"
	myMap[2] = "留步"
	fmt.Println("myMap=",myMap)
	delete(myMap,1)//删除一个key-value
	fmt.Println("myMap=",myMap)

	//字面里创建map
	hqColors := map[string]string{}
	hqColors["red"] = "#FF0000"
	hqColors["green"] = "#00FF00"
	hqColors["blue"] = "#0000FF"
	fmt.Println("hqColors:",hqColors)
}