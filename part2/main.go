package main

import "PublicBlackChain/part2/BLC"

func main() {
	blockchain := BLC.CreatBlockChainWithGenensis()
	blockchain.AddBlockToBlockChain("Send 100RMB To JYY")
	blockchain.AddBlockToBlockChain("Send 200RMB To Freedom")
	blockchain.AddBlockToBlockChain("Send 300RMB To Hope")
	//blockchain.PrintChain()

	blockchain.PrintChainIterator()
}
