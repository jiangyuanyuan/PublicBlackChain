package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器结构体
type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchainIterator *BlockChainIterator) Next() (block *Block) {

	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)

			block = DeSerializeBlock(currentBlockBytes)

			blockchainIterator.CurrentHash = block.PreHash

		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
