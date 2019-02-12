package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块链中的区块,以结构体形式展开
type Block struct {
	Index int64  //区块编码
	Timestamp int64  //创建时间戳
	PreBlockHash string //钱一个
	Hash string  //该区块的hash值
	Data string    //区块信息
}

//计算哈希
//返回加密字符串
func calculateHash(block Block) string{
    //除了hash
    string := string(block.Index) + string(block.PreBlockHash) + block.Data
    //进行加密
    hashInBytes := sha256.Sum256([]byte(string))
    //使用hex十六进制进行解密
    hashInstring :=hex.EncodeToString(hashInBytes[:])
    return hashInstring
}

//添加区块  该区块依赖上一个区块
func GenerateNewBlock(preBlock Block,data string) Block {
    //index
    newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
    return newBlock
}



//创始区块 初始区块链不用返回数据
func GenerateGenisisBlock() Block{
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Timestamp = time.Now().Unix()
	preBlock.PreBlockHash = ""
	preBlock.Hash = calculateHash(preBlock)
	//生成区块
	return GenerateNewBlock(preBlock, "Genesis")
}
