package core

import (
	"encoding/gob"
	"io"
)

// Encoder is an interface that wraps the Encode method.
type Encoder[T any] interface {
	Encode(T) error
}

type TxEncoder struct {
	w io.Writer
}

func NewTxEncoder(w io.Writer) *TxEncoder {
	return &TxEncoder{
		w: w,
	}
}

func (e *TxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(e.w).Encode(tx)
}

type BlockEncoder struct {
	w io.Writer
}

func NewBlockEncoder(w io.Writer) *BlockEncoder {
	return &BlockEncoder{
		w: w,
	}
}

func (enc *BlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(enc.w).Encode(b)
}
