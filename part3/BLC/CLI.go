package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func (cli *CLI) Send(from []string, to []string, acount []string) {
	GetBlockChainObj().MineNewBlock(from, to, acount)
}

func (cli *CLI) GetBalance(from string) {
	GetBlockChainObj().UnUTXOs(from)
}

func (cli *CLI) PrintChain() {
	GetBlockChainObj().PrintChainIterator()
}

func (cli *CLI) CreatBlockChainWithGenensis(data string) {
	CreatBlockChainWithGenensisCLI(data)
}

func (cli *CLI) RUN() {

	isValidArgs()
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	createChainCmd := flag.NewFlagSet("create", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("get", flag.ExitOnError)

	flagFromData := sendBlockCmd.String("from", "", "输入地址")
	flagToData := sendBlockCmd.String("to", "", "输出地址")
	flagAcountData := sendBlockCmd.String("d", "交易数据", "交易金额")
	flagCreateGenensisData := createChainCmd.String("a", "Genensis Block ...", "创世区块地址值")
	flagGetBalanceData := getBalanceCmd.String("a", "", "")
	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "create":
		err := createChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "get":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if sendBlockCmd.Parsed() {
		if *flagFromData == "" || *flagToData == "" || *flagAcountData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagFromData)
		fmt.Println(*flagToData)
		fmt.Println(*flagAcountData)
		cli.Send(JSONToArray(*flagFromData), JSONToArray(*flagToData), JSONToArray(*flagAcountData))
	}

	if printChainCmd.Parsed() {
		fmt.Println("输出所以区块数据")
		cli.PrintChain()
	}

	if createChainCmd.Parsed() {
		if *flagCreateGenensisData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagCreateGenensisData)
		cli.CreatBlockChainWithGenensis(*flagCreateGenensisData)
	}

	if getBalanceCmd.Parsed() {
		if *flagGetBalanceData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagGetBalanceData)
		cli.GetBalance(*flagGetBalanceData)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tsend -from Address -to Address -d Acount")
	fmt.Println("\tcreate -a Address - 创世区块地址值")
	fmt.Println("\tget -a Address - 获取地址值的余额")
	fmt.Println("\tprint - 打印所有区块信息")

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
