package base

import (
	"encoding/binary"
	"go/types"
	"unicode/utf16"
)

type UTF16StringDecoder struct {
	EntityID  string
	Offset    uint32
	Length    uint32
	Endianess Endianess
}

func (t *UTF16StringDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *UTF16StringDecoder) GetEntityType() types.BasicKind {
	return types.String
}

func (t *UTF16StringDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+t.Length > uint32(len(buf)) {
		return nil
	}

	s := buf[base+t.Offset : base+t.Offset+t.Length]

	var utfCodepoints []uint16
	for i := 0; i < int(t.Length/2); i++ {
		if s[i*2] == 0 && s[i*2+1] == 0 {
			break
		}
		if t.Endianess == LittleEndian {
			utfCodepoints = append(utfCodepoints, binary.LittleEndian.Uint16(s[i*2:(i+1)*2]))
		} else {
			utfCodepoints = append(utfCodepoints, binary.BigEndian.Uint16(s[i*2:(i+1)*2]))
		}

	}

	return string(utf16.Decode(utfCodepoints[:]))

}
