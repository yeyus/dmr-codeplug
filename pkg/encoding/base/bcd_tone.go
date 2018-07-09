package base

import (
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type BCDToneDecoder struct {
	EntityID string
	Offset   uint32
}

func (t *BCDToneDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *BCDToneDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *BCDToneDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+2 > uint32(len(buf)) {
		return nil
	}

	s := buf[base+t.Offset : base+t.Offset+2]

	tone := tytera.Tone{}

	// unset
	if s[0] == 0xFF && s[1] == 0xFF {
		return tone
	}

	firstNib := (s[1] & 0xF0) >> 4
	switch {
	case firstNib >= 0 && firstNib <= 7:
		tone.Type = tytera.ToneSystem_CTCSS
		tone.Frequency = (uint32((s[1]&0xF0)>>4) * 1000) +
			(uint32(s[1]&0x0F) * 100) +
			(uint32((s[0]&0xF0)>>4) * 10) +
			uint32(s[0]&0x0F)
	case firstNib == 8:
		tone.Type = tytera.ToneSystem_DCS_NORMAL
		tone.Code = (uint32(s[1]&0x0F) * 100) +
			(uint32((s[0]&0xF0)>>4) * 10) +
			uint32(s[0]&0x0F)
	case firstNib == 0xC:
		tone.Type = tytera.ToneSystem_DCS_INVERTED
		tone.Code = (uint32(s[1]&0x0F) * 100) +
			(uint32((s[0]&0xF0)>>4) * 10) +
			uint32(s[0]&0x0F)
	}

	return tone
}
