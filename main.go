package main

import (
	"PublicBlackChain/BLC"
)

func main() {

	blockchain := BLC.CreatBlockChainWithGenensis()

	blockchain.AddBlockToBlockChain("Send 100RMB To JYY", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("Send 200RMB To Freedom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockChain("Send 300RMB To Hope", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
}
