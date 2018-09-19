package hqstdlib

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)

var configs = make(map[string]string,10)
//开始解析文件
func HqParseFile(filename string)  {

	file,err := os.Open(filename)
	if err != nil {

		return
	}
	defer file.Close()

	//buf,err := ioutil.ReadAll(file)
	//if err != nil {
	//	return
	//}
	//fmt.Println("ccc==",string(buf))
	reader := bufio.NewReader(file)

	for {

		buf,isPrefix,err  := reader.ReadLine()
		fmt.Println("isPrefix==",isPrefix)

		if err == io.EOF{
			break
		}
		parseLine(string(buf))
	}
	fmt.Println(configs)
	fmt.Println("host",configs["host"])
}
func parseLine(line string)  {
	//fmt.Println(line)

	lineStr := strings.TrimSpace(line)
	if len(lineStr) >0 {
		if lineStr[0] == '#' || lineStr[0] == '\n' || lineStr[0] == ';' {
			return
		}
		kvSlice := strings.Split(line,"=")
		configs[kvSlice[0]] = kvSlice[1]
	}


}