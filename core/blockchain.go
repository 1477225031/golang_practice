package core

//处理区块链
//无法删除,无法修改
//可以查看,可以新增
//定义数据的默认结构
type BlockChain struct {
	Blocks []*Block
}
