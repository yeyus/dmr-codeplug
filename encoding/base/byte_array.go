package base

import (
	"go/types"
)

type ByteArrayDecoder struct {
	EntityID  string
	Offset    uint32
	Length    uint32
	Endianess Endianess
}

func (t *ByteArrayDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *ByteArrayDecoder) GetEntityType() types.BasicKind {
	return types.Byte
}

func (t *ByteArrayDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset+t.Length > uint32(len(buf)) {
		return nil
	}

	if t.Endianess == LittleEndian {
		return invert(buf[base+t.Offset : base+t.Offset+t.Length])
	} else {
		return buf[base+t.Offset : base+t.Offset+t.Length]
	}

}

func invert(input []byte) []byte {
	for i := 0; i < len(input)/2; i++ {
		a := input[i]
		b := input[len(input)-i-1]
		input[i] = b
		input[len(input)-i-1] = a
	}

	return input
}
