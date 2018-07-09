package base

import (
	"go/types"
)

type BCDDecoder struct {
	EntityID  string
	Offset    uint32
	Length    uint32
	Endianess Endianess
}

func (t *BCDDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *BCDDecoder) GetEntityType() types.BasicKind {
	return types.Uint64
}

func (t *BCDDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+t.Length > uint32(len(buf)) {
		return nil
	}

	s := buf[base+t.Offset : base+t.Offset+t.Length]

	var value uint64 = 0
	if t.Endianess == LittleEndian {
		for i := 0; uint32(i) < t.Length; i++ {
			value += uint64((s[i] >> 4) & 0xF)
			value *= 10
			value += uint64(s[i] & 0xF)
			if i != int(t.Length-1) {
				value *= 10
			}
		}
	} else {
		for i := int(t.Length) - 1; i >= 0; i-- {
			value += uint64((s[i] >> 4) & 0xF)
			value *= 10
			value += uint64(s[i] & 0xF)
			if i != 0 {
				value *= 10
			}
		}
	}

	return value
}
