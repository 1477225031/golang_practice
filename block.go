package demochain

import (
	"crypto/sha256"
	"encoding/hex"
)

//create a project
//区块链中的区块,以结构体形式展开

type Block struct {
	Index int64  //区块编码
	Hash string  //该区块的hash值
	PreHash string //钱一个
	Data string    //区块信息
}

//计算哈希
//返回加密字符串
func calculateHash(block Block) string{
    //除了hash
    string := string(block.Index) + string(block.PreHash) + block.Data
    //进行加密
    hashInBytes := sha256.Sum256([]byte(string))
    //使用hex十六进制进行解密
    hashInstring :=hex.EncodeToString(hashInBytes[:])
    return hashInstring
}