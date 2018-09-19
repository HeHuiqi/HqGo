package hqstdlib

import (
	"html/template"
	"fmt"
	"os"
	"net/http"
)

type Person struct {
	Id      int
	Name    string
	Country string
	Age int
}

func HqTemplateUse()  {

	/*

		1.创建模板对象
		2.加载模板文件或字符串
		3.执行渲染模板
		*/

	//hqTplStr :="Hello {{ . }}"
	hqTplStr :=` {{  .Name|hqFunc }}`



	tmpl := template.New("ppp")

	fmt.Println("模版名称：",tmpl.Name())

	//函数的注册必须在解析之前
	tmpl.Funcs(template.FuncMap{"hqFunc":hqTplFunc})
	//解析模版字符串 Hello {{.Name}}
	tmpl.Parse(hqTplStr)

	//插入数据
	user := Person{Id: 1001, Name: "小哥", Country: "China"}
	tmpl.Execute(os.Stdout, user)




	//模板函数
	t6 := template.New("test");
	//注册模板函数
	t6.Funcs(template.FuncMap{"test1": test1});
	t6.Funcs(template.FuncMap{"test2": test2});
	// {{函数名}}输出函数返回值
	// {{函数名 参数1 参数2}}
	// {{.字段名|函数名}} 以字段的值作为函数的参数
	t6, _ = t6.Parse(`{{test1}} {{test2 "test2参数"}}`)
	t6.Execute(os.Stdout, Person{12, "test", "214214", 1});


}
func test2(msg string) string {
	fmt.Println("test2==",msg)
	return ""
}
func test1() string {
	fmt.Println("test1")

	return ""
}
func hqTplFunc(args ...interface{}) string {
	//这里只考虑一个参数的情况
	var str string = ""
	if s, ok := args[0].(string); ok {
		str = "*****" + s + "*****"
	} else {
		str = "Nothing"
	}
	return str
}

func HqServerParseTPL()  {
	/*

    1.创建模板对象
    2.加载模板文件或字符串
    3.执行渲染模板
	*/

	//http服务
	http.HandleFunc("/",home)
	http.ListenAndServe("127.0.0.1:8080",nil)
}

func home(w http.ResponseWriter,r *http.Request)  {

	//加载模版文件,加载多个模版文件，模版的名字位第一个模版的名字
	tpl,err := template.ParseFiles("./static/templates/layout.html","static/templates/index.html")
	fmt.Println("HQ:模版名称：",tpl.Name())

	if err != nil {

		fmt.Println(err)
	}
	user := Person{Id: 1001, Name: "小哥模版", Country: "China"}
	user.Age=10
	//执行数据解析
	//tpl.Execute(w,user)
	tpl.ExecuteTemplate(w,"layout",user)
}