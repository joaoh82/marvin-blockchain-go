package core

import "github.com/joaoh82/marvinblockchain/proto"

type HeaderList struct {
	headers []*proto.Header
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

// Add adds a header to the list
func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

// Get returns the header at the given index. The index is 0-based and is also the height of the header
func (list *HeaderList) Get(index int) *proto.Header {
	if index > list.Height() {
		panic("index too high!")
	}
	return list.headers[index]
}

// Last returns the last header in the list. The last header can be used to get the hash of the last block and create a new block.
func (list *HeaderList) Last() *proto.Header {
	return list.Get(list.Height())
}

// Height returns the height of the list. The height is the index of the last header
func (list *HeaderList) Height() int {
	return list.Len() - 1
}

// Len returns the number of headers in the list. Len() - 1 is the height of the list
func (list *HeaderList) Len() int {
	return len(list.headers)
}
