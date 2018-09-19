package hqstdlib

import (
	"log"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL string `json:"unescapeUrl"`
		URL string `json:"url"`
		VisibleURL string `json:"visibleUrl"`
		Title string `json:"title"`
		CacheURL string `json:"cacheUrl"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content string `json:"content"`

	}
	
	gResponse struct {
		ResponseData struct{
			Result[]gResult `json:"results"`
			
		} `json:"responseData"`
	}
)

func HqDecoderJson()  {
	/*
	uri := "file:///Users/hqmac/Desktop/hq/go/HqGoroutine/hqdata/data.json"
	resp,err := http.Get(uri)
	if err !=nil {
		log.Println("ERROR:",err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&gr)

	*/

	jsonPath := "./hqdata/data.json"
	file, openErr := os.OpenFile(jsonPath,os.O_RDWR,0666)
	if openErr != nil {
		log.Println("openErr:",openErr)
		return
	}
	//开始解析
	var gr gResponse
	openErr = json.NewDecoder(file).Decode(&gr)
	if openErr != nil {
		log.Println("ERROR:",openErr)
		return
	}
	fmt.Println(gr.ResponseData)
}

type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct{
		Home string `json:"home"`
		Cell string `json:"cell"`

	} `json:"contact"`
} 
func HqDecoderJson1()  {

	var jsonStr  = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
			"home":"415.566.99966",
			"cell":"1335666"
		}
	}`

	var contact Contact
	//这里解析json字符串要转为[]byte切片
	err := json.Unmarshal([]byte(jsonStr),&contact)
	if err != nil {
		log.Println("jsonERROR1:",err)
	}
	fmt.Println(contact)


	var jsonDataMap map[string] interface{}

	//第一种解析
	err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonDataMap)

	if err != nil {
		fmt.Println("jsonERROR2:",err)
		return
	}
	fmt.Println("json.NewDecoder==",jsonDataMap)

	//第二种解析
	err = json.Unmarshal([]byte(jsonStr),&jsonDataMap)
	if err != nil {
		log.Println("jsonERROR:",err)
		return
	}
	fmt.Println("json.Unmarshal==",jsonDataMap)
	//将内部的对象类型强制转为map类型然后取值
	fmt.Println("cell:",jsonDataMap["contact"].(map[string]interface{})["cell"])


}
func HqEncoderJson()  {
	c := make(map[string]interface{})
	c["name"]="哈哈"
	c["title"]="笑话"
	c["contact"]= map[string]interface{}{
		"home":"上海",
		"call":"123125",
	}
	//jsonData,err := json.Marshal(c)
	jsonData,err := json.MarshalIndent(c,"","  ")//格式化输出
	if err != nil {
		log.Println("encoderError",err)
		return
	}
	fmt.Println("json.MarshalIndent2:",string(jsonData))

}