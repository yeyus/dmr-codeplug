package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type EmergencySystemEntry struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.EmergencySystemEntry
}

func GetEmergencySystemEntry() EmergencySystemEntry {
	m := EmergencySystemEntry{
		EntityID: "com.tytera.emergency.entries[%d]",
		Entry:    tytera.EmergencySystemEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.UTF16StringDecoder{
			EntityID:  "systemName",
			Offset:    0,
			Length:    32,
			Endianess: base.LittleEndian,
		},
		&base.BitMaskDecoder{
			EntityID: "alarmType",
			Offset:   32,
			BitMask:  0x03,
		},
		&base.BitMaskDecoder{
			EntityID: "alarmMode",
			Offset:   32,
			BitMask:  0x30,
		},
		&base.ByteDecoder{
			EntityID: "impoliteRetries",
			Offset:   33,
		},
		&base.ByteDecoder{
			EntityID: "politeRetries",
			Offset:   34,
		},
		&base.ByteDecoder{
			EntityID: "hotMicDuration",
			Offset:   35,
		},
		&base.Uint64Decoder{
			EntityID:  "revertChannel",
			Offset:    36,
			Length:    2,
			Endianess: base.BigEndian,
		},
	}

	return m
}

func (t EmergencySystemEntry) GetEntityID() string {
	return t.EntityID
}

func (t EmergencySystemEntry) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t EmergencySystemEntry) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *EmergencySystemEntry) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "systemName":
		s := d.Decode(buf, base).(string)
		t.Entry.SystemName = string(s)
	case "alarmType":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AlarmType = tytera.AlarmType(uint32(s))
	case "alarmMode":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AlarmMode = tytera.AlarmMode(uint32(s))
	case "impoliteRetries":
		s := d.Decode(buf, base).(uint8)
		t.Entry.ImpoliteRetries = uint32(s)
	case "politeRetries":
		s := d.Decode(buf, base).(uint8)
		t.Entry.PoliteRetries = uint32(s)
	case "hotMicDuration":
		s := d.Decode(buf, base).(uint8)
		t.Entry.HotMicDuration = uint32(s)
	case "revertChannel":
		s := d.Decode(buf, base).(uint64)
		t.Entry.RevertChannel = uint32(s)
	}
}
