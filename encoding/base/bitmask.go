package base

import (
	"go/types"
)

type BitMaskDecoder struct {
	EntityID string
	Offset   uint32
	BitMask  uint8
}

func (t *BitMaskDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *BitMaskDecoder) GetEntityType() types.BasicKind {
	return types.Uint8
}

func (t *BitMaskDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+1 > uint32(len(buf)) {
		return nil
	}

	var leadingzeros uint8 = 0
	var i uint8
	for i = 0; i < 8; i++ {
		if (t.BitMask>>i)&1 == 0 {
			leadingzeros++
		}
	}

	return (buf[base+t.Offset : base+t.Offset+1][0] & t.BitMask) >> leadingzeros
}
