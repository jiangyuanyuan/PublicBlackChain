package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Transaction struct {
	//交易Hash
	TxHash []byte

	//输入
	Vins []*TXInput

	//输出
	Vouts []*TXOutput
}

//创建Transation区分2种情况
//1、创世区块的 input为空的
func NewCoinBaseTransaction(addr string) *Transaction {
	txInput := &TXInput{[]byte{}, -1, "Genensis Data"}
	txOutput := &TXOutput{"10", addr}
	txCoinBase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinBase.TxHash = HashTransaction(txCoinBase)
	return txCoinBase
}

//2、正常交易
func NewNormalTransaction(from string, to string, value string) *Transaction {

	//转账命令：./cli send -from '["freedom"]' -to '["hope"]' -d '["10"]'

	//1、form这个address所有的未话费交易输出的Transaction
	unSpentTx := UnSpentTransationsWithAddress(from)
	fmt.Println(unSpentTx)

	var txInputs []*TXInput
	var txOutputs []*TXOutput
	txInput := &TXInput{[]byte("0000000000000012132131"), -1, from}
	txInputs = append(txInputs, txInput)
	txOutput := &TXOutput{value, to}
	txOutputs = append(txOutputs, txOutput)
	tx := &Transaction{[]byte{}, txInputs, txOutputs}
	tx.TxHash = HashTransaction(tx)
	return tx
}

//交易生成hash
func HashTransaction(tx *Transaction) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	return hash[:]
}
