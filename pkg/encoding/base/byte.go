package base

import (
	"go/types"
)

type ByteDecoder struct {
	EntityID  string
	Offset    uint32
	Endianess Endianess
}

func (t *ByteDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *ByteDecoder) GetEntityType() types.BasicKind {
	return types.Uint8
}

func (t *ByteDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+1 > uint32(len(buf)) {
		return nil
	}

	return buf[base+t.Offset : base+t.Offset+1][0]
}
