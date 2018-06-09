package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type MessagePresets struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Messages tytera.MessagePresets
}

func GetMessagePresets() MessagePresets {
	m := MessagePresets{
		EntityID: "com.tytera.messages",
		Base:     0x2280,
		Length:   0x3840,
		Messages: tytera.MessagePresets{},
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x120,
			Records:      0x32,
			Decoder: &base.UTF16StringDecoder{
				EntityID:  "com.tytera.messages[%d].message",
				Offset:    0,
				Length:    0x120,
				Endianess: base.LittleEndian,
			},
		},
	}

	return m
}

func (t *MessagePresets) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}

	decoder := t.Decoders[0]
	messages := decoder.Decode(buf, base+t.Base).([]interface{})
	for i, ms := range messages {
		id := fmt.Sprintf(decoder.GetEntityID(), i)
		m[id] = ms.(string)
		t.Messages.Messages = append(t.Messages.Messages, ms.(string))
	}

	return m
}
