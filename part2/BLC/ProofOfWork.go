package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBit = 16 //挖矿难度
type ProofOfWork struct {
	Block  *Block
	target *big.Int //1左位移生成做对比的
}

func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit) //左移  生成对比值

	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	//1、拼接block所以属性
	//2、判断hash是否满足难度
	nonce := 0
	var hashInt big.Int //存储新生成的hash
	var hash [32]byte
	for {
		//准备数据
		dataBytes := pow.prepareData(nonce)
		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		//hash 存储hashInt
		hashInt.SetBytes(hash[:])

		//判断hashInt 是否小于target
		if pow.target.Cmp(&hashInt) == 1 {
			fmt.Println()
			break
		}

		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	//5、拼接
	blockBytes := bytes.Join(
		[][]byte{IntToHex(int64(pow.Block.Height)),
			pow.Block.PreHash,
			pow.Block.Data,
			IntToHex(int64(pow.Block.TimesTamp)),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce))},
		[]byte{})
	// 生成hash
	hash := sha256.Sum256(blockBytes)
	return hash[:]
}
