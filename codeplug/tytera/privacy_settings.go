package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type PrivacySettings struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Privacy  tytera.PrivacySettings
}

func GetPrivacySettings() PrivacySettings {
	m := PrivacySettings{
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

	return m
}

func (t *PrivacySettings) Decode(buf []byte, base uint32) map[string]string {
	m := map[string]string{}

	// decode enhanced privacy
	enhancedDecoder := t.Decoders[0]
	enhancedKeys := enhancedDecoder.Decode(buf, base+t.Base).([]interface{})
	t.Privacy.EnhancedKeys = make([][]byte, 8)
	for i, key := range enhancedKeys {
		k := fmt.Sprintf(enhancedDecoder.GetEntityID(), i)
		m[k] = fmt.Sprintf("%X", key)
		t.Privacy.EnhancedKeys[i] = key.([]byte)
	}

	basicDecoder := t.Decoders[1]
	basicKeys := basicDecoder.Decode(buf, base+t.Base).([]interface{})
	t.Privacy.BasicKeys = make([][]byte, 16)
	for i, key := range basicKeys {
		k := fmt.Sprintf(basicDecoder.GetEntityID(), i)
		m[k] = fmt.Sprintf("%X", key)
		t.Privacy.BasicKeys[i] = key.([]byte)
	}

	return m
}
