package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchanIterator *BlockChainIterator) Next() (block *Block) {

	err := blockchanIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			currentBlockBytes := b.Get(blockchanIterator.CurrentHash)

			block = DeSerializeBlock(currentBlockBytes)

			blockchanIterator.CurrentHash = block.PreHash

		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
