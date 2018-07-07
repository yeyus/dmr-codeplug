package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type DTMFGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	DTMF     tytera.DTMFSettings
}

func GetDTMFGroup() DTMFGroup {
	b := DTMFGroup{
		EntityID: "com.tytera.dtmf",
		Base:     0x30100,
		Length:   0x3C0,
		DTMF:     tytera.DTMFSettings{},
	}

	b.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x30,
			Records:      4,
			Decoder:      GetDTMFSystemEntryDecoder(),
		},
		&encoding.RepeatedDecoder{
			Offset:       0xC0,
			RecordLength: 0x10,
			Records:      32,
			Decoder: &base.DTMFStringDecoder{
				EntityID: "com.tytera.dtmf.encode[%d]",
				Offset:   0,
			},
			/*Predicate: func(a interface{}) bool {
				str := a.(string)
				return str != ""
			},*/
		},
		&encoding.RepeatedDecoder{
			Offset:       0x2C0,
			RecordLength: 24,
			Records:      8,
			Decoder:      GetDTMFDecodeEntryDecoder(),
		},
	}

	return b
}

func (t *DTMFGroup) Decode(buf []byte, base uint32) map[string]string {
	m := map[string]string{}

	systemDecoder := t.Decoders[0]
	systems := systemDecoder.Decode(buf, base+t.Base).([]interface{})
	var systemsOut []*tytera.DTMFSystemEntry
	for i, system := range systems {
		entry := system.(tytera.DTMFSystemEntry)
		id := fmt.Sprintf(systemDecoder.GetEntityID(), i)
		m[id] = fmt.Sprintf("%+v", entry)
		systemsOut = append(systemsOut, &entry)
	}
	t.DTMF.Systems = systemsOut

	encodeDecoder := t.Decoders[1]
	encodes := encodeDecoder.Decode(buf, base+t.Base).([]interface{})
	for i, encode := range encodes {
		entry := encode.(string)
		id := fmt.Sprintf(encodeDecoder.GetEntityID(), i)
		m[id] = entry
		t.DTMF.Encodes = append(t.DTMF.Encodes, entry)
	}

	decodeDecoder := t.Decoders[2]
	decodes := decodeDecoder.Decode(buf, base+t.Base).([]interface{})
	var decodesOut []*tytera.DTMFDecodeEntry
	for i, decode := range decodes {
		entry := decode.(tytera.DTMFDecodeEntry)
		id := fmt.Sprintf(encodeDecoder.GetEntityID(), i)
		m[id] = fmt.Sprintf("%+v", entry)
		decodesOut = append(decodesOut, &entry)
	}
	t.DTMF.Decodes = decodesOut

	return m
}
