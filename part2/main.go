package main

import "PublicBlackChain/part2/BLC"

func main() {

	//1、迭代测试
	//blockchain := BLC.CreatBlockChainWithGenensis()
	//blockchain.AddBlockToBlockChain("Send 100RMB To JYY")
	//blockchain.AddBlockToBlockChain("Send 200RMB To Freedom")
	//blockchain.AddBlockToBlockChain("Send 300RMB To Hope")
	////blockchain.PrintChain()
	//
	//blockchain.PrintChainIterator()

	//2、终端工具
	blockchain := BLC.CreatBlockChainWithGenensis()
	CLI := BLC.CLI{blockchain}
	CLI.RUN()

}
