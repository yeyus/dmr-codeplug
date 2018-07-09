package tytera

import (
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type DTMFSystemEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.DTMFSystemEntry
}

func GetDTMFSystemEntryDecoder() DTMFSystemEntryDecoder {
	d := DTMFSystemEntryDecoder{
		EntityID: "com.tytera.dtmf.system[%d]",
		Entry:    tytera.DTMFSystemEntry{},
	}

	d.Decoders = []encoding.Decoder{
		&base.ByteDecoder{
			EntityID: "firstDigitDelay",
			Offset:   0,
		},
		&base.ByteDecoder{
			EntityID: "firstDigitTime",
			Offset:   1,
		},
		&base.ByteDecoder{
			EntityID: "digitDurationTime",
			Offset:   2,
		},
		&base.ByteDecoder{
			EntityID: "digitIntervalTime",
			Offset:   3,
		},
		&base.ByteDecoder{
			EntityID: "starHashDigitTime",
			Offset:   4,
		},
		&base.ByteDecoder{
			EntityID: "dKeyAssignment",
			Offset:   5,
		},
		&base.ByteDecoder{
			EntityID: "nextSequence",
			Offset:   6,
		},
		&base.ByteDecoder{
			EntityID: "autoResetTime",
			Offset:   7,
		},
		&base.BitDecoder{
			EntityID:  "sideTone",
			Offset:    8,
			BitOffset: 0,
		},
		&base.BitMaskDecoder{
			EntityID: "pttId",
			Offset:   8,
			BitMask:  0x0C,
		},
		&base.BitMaskDecoder{
			EntityID: "groupCode",
			Offset:   8,
			BitMask:  0xF0,
		},
		&base.DTMFStringDecoder{
			EntityID: "keyUpEncode",
			Offset:   16,
		},
		&base.DTMFStringDecoder{
			EntityID: "keyDownEncode",
			Offset:   32,
		},
	}

	return d
}

func (t DTMFSystemEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t DTMFSystemEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t DTMFSystemEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *DTMFSystemEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "firstDigitDelay":
		s := d.Decode(buf, base).(uint8)
		t.Entry.FirstDigitDelay = uint32(s) * 10
	case "firstDigitTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.FirstDigitTime = uint32(s) * 10
	case "digitDurationTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.DigitDurationTime = uint32(s) * 10
	case "digitIntervalTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.DigitIntervalTime = uint32(s) * 10
	case "starHashDigitTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.StarHashDigitTime = uint32(s) * 10
	case "dKeyAssignment":
		s := d.Decode(buf, base).(uint8)
		t.Entry.DKeyAssignment = uint32(s) * 10
	case "nextSequence":
		s := d.Decode(buf, base).(uint8)
		t.Entry.NextSequence = uint32(s)
	case "autoResetTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AutoResetTime = uint32(s)
	case "sideTone":
		s := d.Decode(buf, base).(bool)
		t.Entry.SideTone = s
	case "pttId":
		s := d.Decode(buf, base).(uint8)
		t.Entry.PttId = tytera.PttIdType(int(s))
	case "groupCode":
		s := d.Decode(buf, base).(uint8)
		t.Entry.GroupCode = tytera.GroupCodeType(int(s))
	case "keyUpEncode":
		s := d.Decode(buf, base).(string)
		t.Entry.KeyUpEncode = s
	case "keyDownEncode":
		s := d.Decode(buf, base).(string)
		t.Entry.KeyDownEncode = s
	}
}
