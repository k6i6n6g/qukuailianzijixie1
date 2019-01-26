package main

import (
	"bytes"
	"fmt"
)

func main(){
	ByteArray:=[][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
	}
	result:=bytes.Join(ByteArray,[]byte(""))
	fmt.Printf("result:%s\n",string(result))
}