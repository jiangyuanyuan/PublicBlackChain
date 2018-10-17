package BLC

type Transaction struct {
	//交易Hash
	TxHash []byte

	//输入
	Vins []*TXInput

	//输出
	Vouts []*TXOutput
}
