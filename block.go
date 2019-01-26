package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

    //创世神块
const genesisInfo="The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
type Block struct {
	//版本号
	Version uint64
	//上一个哈希
	PrevHash  []byte
	//梅克尔根
	MerkleRoot  []byte
	//时间戳
	TimeStamp  uint64
	//难度值
	Difficulty  uint64
	//随机数
	Nonce   uint64
	//当前哈希
	Hash  []byte
	//区块数据
	Data  []byte
}
//创建创世块
func GenesisBlock(data string,prevBlockHash[]byte)*Block{
	return  NewBlock(prevBlockHash,data)
}
//生成哈希值
func (block *Block)SetHash(){
	//储存拼接好的数据,最后作为sha256函数的参数
	var blockByteInfo []byte
	/*//拼接当前区块数据
	blockByteInfo=append(blockByteInfo,block.PrevHash...)
	//blockByteInfo=append(blockByteInfo,b.Hash...)
	blockByteInfo=append(blockByteInfo,block.Data...)
	blockByteInfo=append(blockByteInfo,block.MerkleRoot...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Version)...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.TimeStamp)...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Difficulty)...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Nonce)...)*/
	//上面的太复杂了 所以用join来拼接
	tmp:=[][]byte{
		block.MerkleRoot,
		block.PrevHash,
		block.Data,
		uint64ToByte(block.Nonce),
		uint64ToByte(block.Difficulty),
		uint64ToByte(block.TimeStamp),
		uint64ToByte(block.Version),
	}
	blockByteInfo=bytes.Join(tmp,[]byte(""))
	//进行哈希256运算
	hash:=sha256.Sum256(blockByteInfo)
	//把哈希添加到区块hash字段中
	block.Hash=hash[:]
}
//辅助生成哈希的    将字符串形式变成二进制
func uint64ToByte(num uint64)[]byte{
	//buffer 意思为缓冲器
	var buffer bytes.Buffer
	//将数据以二进制形式保存到buffer中    binary序列化     大端位就是从小到大排序
	err:=binary.Write(&buffer ,binary.BigEndian/*大端对齐*/,num)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}
//创建区块     传入哈希 与区块数据
func NewBlock(prevBlockHash []byte, data string)*Block{
	block:=Block{
		Version:00,
		PrevHash: prevBlockHash,
		MerkleRoot:[]byte{},//v4版本再讲
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty: 100,
		Nonce:100,
		//Hash:[]byte{},
		Data :[]byte(data),
	}
	//调用hash值
	block.SetHash()
	return  &block
}
//添加区块
func (bc *BlockChain)AddBlock(data string){
	//bc,Blocks 为一条链中区块的长度
	blockLen:=len(bc.Blocks)
	//获取最后一个区块的哈希值
	lastBlock:=bc.Blocks[blockLen-1]
	//最后一个区块的哈希值是新区块的前哈希
	prevBlockHash:=lastBlock.Hash
	//变成一个新的qukuai
	block:=NewBlock(prevBlockHash,data)
	//向bc中添加新的区块
	bc.Blocks=append(bc.Blocks,block)

}