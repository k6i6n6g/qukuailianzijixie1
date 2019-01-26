package main

import (
	"fmt"
)

func main(){
	//block:=NewBlock([]byte{0x0000000000000000},genesisInfo)
	bc:=NewBlockChain()
	bc.AddBlock("坚持!")
	bc.AddBlock("进步!")
	bc.AddBlock("胜利!")
	//bc.Blocks这是一条链中的一个区块
	for _,block:=range bc.Blocks{
		fmt.Printf("PrevHash:%x\n",block.PrevHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("Data:%x\n",block.Data)
	}
}
