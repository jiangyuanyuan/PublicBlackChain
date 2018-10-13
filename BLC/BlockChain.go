package BLC

type BlockChain struct {
	Blocks []*Block
}

//创建创世区块链
func CreatBlockChainWithGenensis() *BlockChain {
	genensisBlock := CreateGenensisBlock("Genenis Block ...")

	return &BlockChain{[]*Block{genensisBlock}}
}

func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, preHash []byte) {
	//创建新区块
	block := NewBlock(height, preHash, data)
	//添加到区块链中
	blc.Blocks = append(blc.Blocks, block)
}
