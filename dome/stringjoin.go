package main

import (
	"fmt"
	"strings"
)

func main()  {
	//strings为包名
	strArray:=[]string{"hello","world","itcast"}
	// string.jion就是把strArray用+链接
	result:=strings.Join(strArray,"+")
	fmt.Printf("result:%s\n",result)
	result=strings.Join(strArray,"")
	fmt.Printf("result:%s\n",result)
}
