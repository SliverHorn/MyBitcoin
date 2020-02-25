package main

import (
	"fmt"
	"github.com/SliverHorn/MyBitcoin/MyBitcoin1/Core"
)

func main() {
	bc := Core.NewBlockChain()                            // 初始化区块链, 创建第一个区块(创世界区块)
	bc.AddBlock("Send 1 SliverHorn1 to SliverHorn2")      // 加入一个区块,发送一个比特币给SliverHorn2
	bc.AddBlock("Send 2 more SliverHorn1 to SliverHorn2") // 加入一个区块.发送更多比特币给SliverHorn2

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%x\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println()
	}
}
