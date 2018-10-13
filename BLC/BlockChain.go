package BLC

type BlockChain struct {
	Blocks []*Block
}

//创建创世区块链
func CreatBlockChainWithGenensis() *BlockChain {
	genensisBlock := CreateGenensisBlock("Genenis Block ...")

	return &BlockChain{[]*Block{genensisBlock}}
}
