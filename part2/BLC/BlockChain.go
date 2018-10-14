package BLC

import (
	"PublicBlackChain/part1/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

const dbName = "block.db"
const bucketName = "blocks"

type BlockChain struct {
	//Blocks []*Block

	//加入db 持久存储
	Tip []byte   //最新区块的hash
	DB  *bolt.DB //DB
}

//创建创世区块链
func CreatBlockChainWithGenensis() *BlockChain {
	var blockHash []byte
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			b, err = tx.CreateBucket([]byte(bucketName))
			if err != nil {
				log.Panic(err)
			}
		}
		genensisBlock := CreateGenensisBlock("Genenis Block ...")
		err = b.Put([]byte(genensisBlock.Hash), genensisBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		blockHash = genensisBlock.Hash
		err = b.Put([]byte("l"), genensisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	return &BlockChain{blockHash, db}
}

//添加到区块到DB中
func (blc *BlockChain) AddBlockToBlockChain(data string) {
	//创建新区块
	//block := NewBlock(height, preHash, data)
	//添加到区块链中
	//blc.Blocks = append(blc.Blocks, block)
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	blc.DB = db
	//添加到DB中
	err = blc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			_, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				log.Panic(err)
			}
		}
		//取出最
		blockBytes := b.Get(blc.Tip)
		preBlock := BLC.DeSerializeBlock(blockBytes)
		//创建新区块
		block := NewBlock(preBlock.Height+1, preBlock.Hash, data)
		err := b.Put([]byte(block.Hash), block.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), block.Hash)
		blc.Tip = block.Hash
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	defer blc.DB.Close()
}

//遍历所有区块信息
func (blc *BlockChain) PrintChain() {
	var block *Block
	var currentHash []byte = blc.Tip
	var genensisHashBytes = big.NewInt(0)

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	blc.DB = db
	for {
		err := blc.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketName))
			if b != nil {
				blockBytes := b.Get(currentHash)
				block = DeSerializeBlock(blockBytes)
				fmt.Println(block)
			}
			return nil

		})
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PreHash)
		if genensisHashBytes.Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.PreHash

	}
}

//迭代器
func (blc *BlockChain) Iterator() *BlockChainIterator {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &BlockChainIterator{blc.Tip, db}
}

//迭代遍历
func (blc *BlockChain) PrintChainIterator() {
	blockChainIterator := blc.Iterator()
	defer blockChainIterator.DB.Close()
	var hashInt big.Int
	var genensis = big.NewInt(0)
	for {
		block := blockChainIterator.Next()
		hashInt.SetBytes(block.PreHash)
		fmt.Println(block)
		if genensis.Cmp(&hashInt) == 0 {
			break
		}
	}
}
