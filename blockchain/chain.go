package blockchain

import (
	"sync"

	"github.com/DianaLeee/gocoin/db"
	"github.com/DianaLeee/gocoin/utils"
)

var b *blockchain; 
var once sync.Once

func (b * blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b));
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1);
	b.NewestHash = block.Hash;
	b.Height = block.Height;
}

func Blockchain() *blockchain {
	if b == nil { 
		once.Do(func() {
			b = &blockchain{"", 0};
			b.AddBlock("Genesis Block")
		})
	}
	return b;
}