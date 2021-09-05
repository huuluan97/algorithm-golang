package blockchain_go_vietnamese

import (
	"bytes"
)

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	//encoder := gob.NewEncoder(&result)
	//err := encoder.Encode(b)

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	//decoder := gob.NewDecoder(bytes.NewReader(d))
	//err := decoder.Decode(&block)

	return &block
}