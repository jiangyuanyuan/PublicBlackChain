package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	//区块高度
	Height int64
	//上一个区块的hash
	PreHash []byte
	//交易数据
	Data []byte
	//区块Hash
	Hash []byte
	//挖矿所需的
	//Nonce int64
	//交易时间戳
	TimesTamp int64
}

//创建创世区块
func CreateGenensisBlock(data string) *Block {
	return NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, data)

}

//创建一个区块
func NewBlock(Height int64, PreHash []byte, data string) *Block {
	block := &Block{Height, PreHash, []byte(data), nil, time.Now().Unix()}
	block.SetHash()
	fmt.Println(block)
	return block
}

func (block *Block) SetHash() {
	//拼接所以属性生成Hash

	//1、Height 转换 []byte数组
	heighBytes := IntToHex(block.Height)
	//fmt.Println(heighBytes)
	//2、preHash 转换 []byte数组

	//3、data 转换 []byte数组

	//4、TimesTamp 转换 []byte数组
	timeString := strconv.FormatInt(block.TimesTamp, 2)
	timeBytes := []byte(timeString)
	//fmt.Println(timeBytes)
	//5、拼接
	blockBytes := bytes.Join([][]byte{heighBytes, block.PreHash, block.Data, block.Hash, timeBytes}, []byte{})
	// 生成hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]

}
