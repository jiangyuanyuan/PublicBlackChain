package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//区块高度
	Height int64
	//上一个区块的hash
	PreHash []byte
	//交易数据
	txs []*Transaction
	//区块Hash
	Hash []byte

	//交易时间戳
	TimesTamp int64
	//挖矿所需的
	Nonce int64
}

//创建创世区块
func CreateGenensisBlock(txs []*Transaction) *Block {

	return NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, txs)

}

//创建一个区块
func NewBlock(Height int64, PreHash []byte, txs []*Transaction) *Block {
	//创建区块对象
	block := &Block{Height, PreHash, txs, nil, time.Now().Unix(), 0}

	//工作量证明   生成有效Hash、nonce值
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	fmt.Println(block)
	fmt.Println(hash)
	fmt.Println(nonce)
	return block
}

//叠加交易生成[]byte
func (blcok *Block) HashTransactions() []byte {

	var txHashs [][]byte
	var txHash [32]byte
	for _, tx := range blcok.txs {
		txHashs = append(txHashs, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashs, []byte{}))

	return txHash[:]
}

//序列化  便于存储到db中
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化 db中取出的byte 生成block对象
func DeSerializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
