package BLC

type UTXO struct {
	TxHash []byte
	Index  int64
	Output *TXOutput
}
