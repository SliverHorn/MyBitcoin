package Core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block headers (区块结构体)
type Block struct {
	TimeStamp     int64		// 区块链创建时间戳
	Data          []byte	// 区块链包含的数据
	PrevBlockHash []byte	// 前一个区块的哈希值
	Hash          []byte	// 区块自身的哈希值, 用于校验区块数据有效
}

// NewBlock creates and return Block(创建并返回Block)
func NewBlock(Date string, PrevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(Date),
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

// SetHash calculates and sets block hash(计算和设置块散列)
func (b *Block) SetHash()  {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block By SliverHorn 7777", []byte{})
}
