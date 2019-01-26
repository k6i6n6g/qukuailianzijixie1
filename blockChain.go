package main
type BlockChain struct {
//定义一个区块链结合 :定一个[]*Block
	Blocks []*Block
}
//创建区块链
func NewBlockChain()*BlockChain{
	//对这个区块链进行初始化,就是把创世块作为第一个元素添加进去
	genesisBlock:=NewBlock([]byte{byte(0x0000000000000000)}, genesisInfo)
	bc:=BlockChain{[]*Block{genesisBlock}}
	return &bc
}
