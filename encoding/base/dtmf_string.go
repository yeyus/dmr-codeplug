package base

import (
	"bytes"
	"fmt"
	"go/types"
)

const LENGTH int = 16

type DTMFStringDecoder struct {
	EntityID string
	Offset   uint32
}

func (t *DTMFStringDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *DTMFStringDecoder) GetEntityType() types.BasicKind {
	return types.String
}

func (t *DTMFStringDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+uint32(LENGTH) > uint32(len(buf)) {
		return nil
	}

	var buffer bytes.Buffer
	l := buf[base+t.Offset : base+t.Offset+1][0]
	s := buf[base+t.Offset+1 : base+t.Offset+uint32(LENGTH)]

	for i := 0; i < min(int(l), LENGTH*2); i++ {
		var a byte
		if i%2 == 1 {
			a = ((s[i/2] >> 4) & 0xF)
		} else {
			a = (s[i/2] & 0xF)
		}

		buffer.WriteString(dtmfCharConvert(a))
	}

	return buffer.String()
}

func dtmfCharConvert(code byte) (c string) {
	if code >= 0 && code < 10 {
		c = fmt.Sprintf("%d", code)
	} else if code == 10 {
		c = "A"
	} else if code == 11 {
		c = "B"
	} else if code == 12 {
		c = "C"
	} else if code == 13 {
		c = "D"
	} else if code == 14 {
		c = "*"
	} else if code == 15 {
		c = "#"
	} else {
		c = "?"
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
