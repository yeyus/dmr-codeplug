package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type PrivacySettingsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Privacy  tytera.PrivacySettings
}

func GetPrivacySettingsGroup() *PrivacySettingsGroup {
	m := PrivacySettingsGroup{
		EntityID: "com.tytera.privacy",
		Base:     0x5AC0,
		Length:   0xAF,
		Privacy:  tytera.PrivacySettings{},
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 16,
			Records:      8,
			Decoder: &base.ByteArrayDecoder{
				EntityID:  "com.tytera.privacy.enhanced[%d]",
				Offset:    0,
				Length:    16,
				Endianess: base.LittleEndian,
			},
		},
		&encoding.RepeatedDecoder{
			Offset:       0x90,
			RecordLength: 2,
			Records:      16,
			Decoder: &base.ByteArrayDecoder{
				EntityID:  "com.tytera.privacy.basic[%d]",
				Offset:    0,
				Length:    2,
				Endianess: base.LittleEndian,
			},
		},
	}

	return &m
}

func (t *PrivacySettingsGroup) GetEntityID() string {
	return t.EntityID
}

func (t *PrivacySettingsGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *PrivacySettingsGroup) Decode(buf []byte, base uint32) interface{} {
	// decode enhanced privacy
	enhancedDecoder := t.Decoders[0]
	enhancedKeys := enhancedDecoder.Decode(buf, base+t.Base).([]interface{})
	t.Privacy.EnhancedKeys = make([][]byte, 8)
	for i, key := range enhancedKeys {
		t.Privacy.EnhancedKeys[i] = key.([]byte)
	}

	basicDecoder := t.Decoders[1]
	basicKeys := basicDecoder.Decode(buf, base+t.Base).([]interface{})
	t.Privacy.BasicKeys = make([][]byte, 16)
	for i, key := range basicKeys {
		t.Privacy.BasicKeys[i] = key.([]byte)
	}

	return t.Privacy
}
