package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64  //区块编号
	Timestamp    int64  //时间戳
	PreBlockHash string //	上一个区块的哈希值
	Hash         string //当前区块hash

	Data string //区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PreBlockHash + b.Data
	byteInBytes := sha256.Sum256([]byte(blockData))

	return hex.EncodeToString(byteInBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
