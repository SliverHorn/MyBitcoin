package Core

// BlockChain keeps a sequence of Block(区块链的结构体)
type BlockChain struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain(将提供的数据保存为区块链中的一个块)
func (bc *BlockChain) AddBlock(data string) {
	PrevBlock := bc.Blocks[len(bc.Blocks)-1]
	NewBlock := NewBlock(data, PrevBlock.Hash)
	bc.Blocks = append(bc.Blocks, NewBlock)
}

// NewBlockChain creates a new Blockchain with genesis Block (创建一个新的区块链含创世纪块)
func NewBlockChain() *BlockChain {
	return &BlockChain{Blocks:[]*Block{NewGenesisBlock()}}
}
