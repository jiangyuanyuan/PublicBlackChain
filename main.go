package main

import (
	"PublicBlackChain/BLC"
	"fmt"
)

func main() {

	BlockChain := BLC.CreatBlockChainWithGenensis()
	fmt.Println(BlockChain.Blocks[0])
}
