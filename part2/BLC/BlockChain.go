package BLC

import (
	"github.com/boltdb/bolt"
	"log"
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
	return &BlockChain{[]byte("l"), db}
}

func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, preHash []byte) {
	//创建新区块
	//block := NewBlock(height, preHash, data)
	//添加到区块链中
	//blc.Blocks = append(blc.Blocks, block)

	//添加到DB中
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			_, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				log.Panic(err)
			}
		}
		//创建新区块
		block := NewBlock(height, preHash, data)
		err := b.Put([]byte(block.Hash), block.Serialize())
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
