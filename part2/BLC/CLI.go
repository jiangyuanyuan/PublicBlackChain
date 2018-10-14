package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BlockChain *BlockChain
}

func (cli *CLI) AddBlock(data string) {
	cli.BlockChain.AddBlockToBlockChain(data)
}

func (cli *CLI) PrintChain() {
	cli.BlockChain.PrintChainIterator()
}

func (cli *CLI) RUN() {

	isValidArgs()
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	flagAddData := addBlockCmd.String("d", "freedom", "交易数据")
	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddData)
		cli.AddBlock(*flagAddData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("输出所以区块数据")
		cli.PrintChain()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tadd -d DATA - 交易数据")
	fmt.Println("\tprint - 打印信息")
	fmt.Println("Usage:")

}
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
