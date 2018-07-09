package base

import (
	"go/types"
)

type Uint64Decoder struct {
	EntityID string
	Offset   uint32
	// Length can't be bigger than 8 bytes
	Length    uint32
	Endianess Endianess
}

func (t *Uint64Decoder) GetEntityID() string {
	return t.EntityID
}

func (t *Uint64Decoder) GetEntityType() types.BasicKind {
	return types.Uint64
}

func (t *Uint64Decoder) Decode(buf []byte, base uint32) interface{} {
	if t.Length > 8 || base+t.Offset > uint32(len(buf)) || base+t.Offset+t.Length > uint32(len(buf)) {
		return nil
	}

	var result uint64
	if t.Endianess == LittleEndian {
		result = uvarint_le(buf[base+t.Offset : base+t.Offset+t.Length])
	} else {
		result = uvarint_be(buf[base+t.Offset : base+t.Offset+t.Length])
	}

	return result
}

func uvarint_le(buf []byte) (x uint64) {
	for i, b := range buf {
		x = x<<8 + uint64(b)
		if i == 7 {
			return
		}
	}
	return
}

func uvarint_be(buf []byte) (x uint64) {
	for i := len(buf) - 1; i >= 0; i-- {
		x = x<<8 + uint64(buf[i])
		if i == 0 {
			return
		}
	}
	return
}
