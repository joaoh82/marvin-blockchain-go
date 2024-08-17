package core

import (
	"encoding/gob"
	"io"
)

// Decoder is an interface that wraps the Decode method.
type Decoder[T any] interface {
	Decode(T) error
}

type TxDecoder struct {
	r io.Reader
}

func NewTxDecoder(r io.Reader) *TxDecoder {
	return &TxDecoder{
		r: r,
	}
}

func (e *TxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)
}

type BlockDecoder struct {
	r io.Reader
}

func NewBlockDecoder(r io.Reader) *BlockDecoder {
	return &BlockDecoder{
		r: r,
	}
}

func (dec *BlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(dec.r).Decode(b)
}
