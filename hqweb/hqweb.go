package hqweb

import (
	"net/http"
	"io"
	"fmt"
)

func HqNewWebServer()  {


	//简单web服务器
	http.Handle("/",&HqHandler{})
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {

	})
	var addr = "127.0.0.1:8080"
	fmt.Printf("http://%s","127.0.0.1:8080")
	http.ListenAndServe(addr,nil)



	//设置服务器参数
	//s := &http.Server{
	//	Addr:"127.0.0.1:8080",
	//	Handler:&HqHandler{},
	//	ReadTimeout:60*time.Second,
	//	WriteTimeout:60*time.Second,
	//	MaxHeaderBytes:1<<20,
	//}
	//log.Fatalln(s.ListenAndServe())
}



type HqHandler struct {

}
//实现Handler接口方法拦截请求
func (handler *HqHandler)ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	homePage(writer,request)
}
func homePage(writer http.ResponseWriter, request *http.Request)  {

	io.WriteString(writer,"home")
}