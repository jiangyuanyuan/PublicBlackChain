package main

import (
	"PublicBlackChain/part1/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//1、创建创世区块链、添加新区块测试
	//blockchain := BLC.CreatBlockChainWithGenensis()
	//blockchain.AddBlockToBlockChain("Send 100RMB To JYY", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockChain("Send 200RMB To Freedom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockChain("Send 300RMB To Hope", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	//2、有效hash验证测试
	//block:=BLC.NewBlock(1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},"test")
	//proofOfWork:=BLC.NewProofOfWork(block)
	//fmt.Printf("%v",proofOfWork.IsValid())

	//3、序列号 反序列化测试 便于存储到DB
	//block:=BLC.NewBlock(1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,},"test")
	//bytes :=block.Serialize()
	//fmt.Println(bytes)
	//block = BLC.DeSerializeBlock(bytes)
	//fmt.Println(block)

	//4、boltdb测试
	block := BLC.NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "test")
	db, err := bolt.Open("block.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic(err)
			}
		}
		err = b.Put([]byte("l"), block.Serialize())
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			fmt.Println(blockData)
			block := BLC.DeSerializeBlock(blockData)
			fmt.Println(block)
		}
		return nil
	})
}
