package base

import (
	"go/types"
	"strings"
)

type ASCIIStringDecoder struct {
	EntityID string
	Offset   uint32
	Length   uint32
}

func (t *ASCIIStringDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *ASCIIStringDecoder) GetEntityType() types.BasicKind {
	return types.String
}

func (t *ASCIIStringDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+t.Length > uint32(len(buf)) {
		return nil
	}

	s := buf[base+t.Offset : base+t.Offset+t.Length]

	end := strings.Index(string(s), "\x00")
	if end == -1 {
		end = len(s)
	}

	return string(s[:end])
}
