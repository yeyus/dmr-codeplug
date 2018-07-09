package base

import (
	"go/types"
)

type BitDecoder struct {
	EntityID  string
	Offset    uint32
	BitOffset uint8
}

func (t *BitDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *BitDecoder) GetEntityType() types.BasicKind {
	return types.Bool
}

func (t *BitDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) ||
		base+t.Offset+1 > uint32(len(buf)) ||
		t.BitOffset > 7 {
		return nil
	}

	s := buf[base+t.Offset : base+t.Offset+1]

	if (s[0]>>t.BitOffset)&1 == 1 {
		return true
	} else {
		return false
	}
}
