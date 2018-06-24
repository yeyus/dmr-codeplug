package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type MessagePresetsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Messages tytera.MessagePresets
}

func GetMessagePresetsGroup() MessagePresetsGroup {
	m := MessagePresetsGroup{
		EntityID: "com.tytera.messages",
		Base:     0x2280,
		Length:   0x3840,
		Messages: tytera.MessagePresets{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(string)

		return e != ""
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
			Predicate: predicate,
		},
	}

	return m
}

func (t *MessagePresetsGroup) Decode(buf []byte, base uint32) (m map[string]string) {
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
