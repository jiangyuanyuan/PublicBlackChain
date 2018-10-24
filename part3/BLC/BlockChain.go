package BLC

import (
	"PublicBlackChain/part1/BLC"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
)

const dbName = "block.db"
const bucketName = "blocks"

type BlockChain struct {
	//加入db 持久存储
	Tip []byte   //最新区块的hash
	DB  *bolt.DB //DB
}

//创建创世区块链
func CreatBlockChainWithGenensisCLI(data string) {
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
		//创建一个基础交易
		transaction := NewCoinBaseTransaction(data)
		genensisBlock := CreateGenensisBlock([]*Transaction{transaction})
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

}

//转账
func (blc *BlockChain) MineNewBlock(from []string, to []string, acount []string) {
	//通过相关算法建立Transaction数组
	var txs []*Transaction

	value, _ := strconv.Atoi(acount[0])
	//建立一笔转账
	tx := NewNormalTransaction(from[0], to[0], int64(value))

	txs = append(txs, tx)

	//添加到DB中
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			_, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				log.Panic(err)
			}
		}
		//取出最顶的block
		blockBytes := b.Get(blc.Tip)
		preBlock := BLC.DeSerializeBlock(blockBytes)
		//创建新区块
		block := NewBlock(preBlock.Height+1, preBlock.Hash, txs)
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

//获取blockchain对象
func GetBlockChainObj() *BlockChain {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var tipHash []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			tipHash = b.Get([]byte("l"))

		} else {
			fmt.Println("先创建创世区块")
		}
		return nil
	})
	return &BlockChain{tipHash, db}
}

//返回该地址下所有未花费交易输出
func (blc *BlockChain) UnUTXOs(address string) []*UTXO {
	blockChainIterator := blc.Iterator()
	defer blockChainIterator.DB.Close()

	var unUTXOs []*UTXO

	var hashInt big.Int
	var genensis = big.NewInt(0)
	var spentTXOutputs = make(map[string][]int64)
	for {
		block := blockChainIterator.Next()
		hashInt.SetBytes(block.PreHash)
		fmt.Println(block)

		for _, tx := range block.txs {
			if tx.IsConbaseTransaction() == false {
				for _, in := range tx.Vins {
					if in.UnLockWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)

						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)

					}
				}
			}

			for index, out := range tx.Vouts {
				if out.UnLockWithAddress(address) {
					if spentTXOutputs != nil {
						if len(spentTXOutputs) != 0 {
							for txHash, indexArray := range spentTXOutputs {
								for _, i := range indexArray {
									if int64(index) == i && txHash == hex.EncodeToString(tx.TxHash) {
										continue
									} else {
										utxo := &UTXO{tx.TxHash, index, out}
										unUTXOs = append(unUTXOs, utxo)
									}
								}
							}
						} else {
							utxo := &UTXO{tx.TxHash, index, out}
							unUTXOs = append(unUTXOs, utxo)
						}

					}
				}
			}

		}

		if genensis.Cmp(&hashInt) == 0 {
			break
		}
	}
	return unUTXOs
}

func (blc *BlockChain) FindSpendableUTXOS(from string, amount int64) (int64, map[string][]int64) {

	//获取我所以的钱UTXO
	utxos := blc.UnUTXOs(from)

	var value int64
	var unSpendableUTXOS = make(map[string][]int64)
	for _, utxo := range utxos {
		value = value + utxo.Output.Value
		hash := hex.EncodeToString(utxo.TxHash)
		unSpendableUTXOS[hash] = append(unSpendableUTXOS[hash], utxo.Index)

		if value >= amount {
			break
		}
	}

	if value < amount {
		fmt.Println("余额不足")
		os.Exit(1)
	}
	return value, unSpendableUTXOS
}

//查询余额
func (blc *BlockChain) GetBalance(address string) (amount int64) {
	utxos := blc.UnUTXOs(address)
	for _, utxo := range utxos {
		amount = amount + utxo.Output.Value
	}
	return
}

//迭代器
func (blc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{blc.Tip, blc.DB}
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
		//fmt.Println(block)
		fmt.Println("####################################################################################")
		fmt.Println()
		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Printf("Height------>| %d\n", block.Height)
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Printf("PreHash----->| %x\n", block.PreHash)
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Printf("Hash-------->| %x\n", block.Hash)
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Printf("Nonce------->| %d\n", block.Nonce)
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Printf("TimesTamp--->| %d\n", block.TimesTamp)
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Println("*********************************->   Txs   <-**************************************")
		for _, tx := range block.txs {
			fmt.Printf("Hash-------->: %x\n", tx.TxHash)
			for _, in := range tx.Vins {
				fmt.Println(in)
			}
			for _, out := range tx.Vouts {
				fmt.Println(out)
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("####################################################################################")
		if genensis.Cmp(&hashInt) == 0 {
			break
		}
	}
}
