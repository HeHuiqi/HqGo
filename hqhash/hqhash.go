package hqhash

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
)

func HqHash()  {

	str := "你好"
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	result := md5Hash.Sum([]byte(""))
	fmt.Println(result)

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(str))
	result = sha1Hash.Sum([]byte(""))
	fmt.Println(result)


}