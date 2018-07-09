package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type OneTouchAccessEntry struct {
	EntityID            string
	Decoders            []encoding.Decoder
	OneTouchAccessEntry tytera.OneTouchAccessEntry
}

func GetOneTouchAccessEntry() OneTouchAccessEntry {
	m := OneTouchAccessEntry{
		EntityID:            "com.tytera.buttons.onetouch[%d]",
		OneTouchAccessEntry: tytera.OneTouchAccessEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.BitMaskDecoder{
			EntityID: "com.tytera.buttons.onetouch.mode",
			Offset:   0,
			BitMask:  0x30,
		},
		&base.BitMaskDecoder{
			EntityID: "com.tytera.buttons.onetouch.calltype",
			Offset:   0,
			BitMask:  0x03,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.onetouch.messageencode",
			Offset:   1,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.onetouch.call",
			Offset:   2,
		},
	}

	return m
}

func (t OneTouchAccessEntry) GetEntityID() string {
	return t.EntityID
}

func (t OneTouchAccessEntry) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t OneTouchAccessEntry) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.OneTouchAccessEntry
}

func (t *OneTouchAccessEntry) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.buttons.onetouch.mode":
		s := d.Decode(buf, base).(uint8)
		t.OneTouchAccessEntry.Mode = tytera.SystemType(uint32(s))
	case "com.tytera.buttons.onetouch.calltype":
		s := d.Decode(buf, base).(uint8)
		ctype := tytera.OneTouchCallType_NO_CALL_TYPE
		if t.OneTouchAccessEntry.Mode == tytera.SystemType_DIGITAL {
			if s == 0 {
				ctype = tytera.OneTouchCallType_CALL
			} else if s == 1 {
				ctype = tytera.OneTouchCallType_TEXT_MESSAGE
			}
		} else if t.OneTouchAccessEntry.Mode == tytera.SystemType_ANALOG {
			switch s {
			case 0:
				ctype = tytera.OneTouchCallType_DTMF1
			case 1:
				ctype = tytera.OneTouchCallType_DTMF2
			case 2:
				ctype = tytera.OneTouchCallType_DTMF3
			case 3:
				ctype = tytera.OneTouchCallType_DTMF4
			}
		}
		t.OneTouchAccessEntry.CallType = ctype
	case "com.tytera.buttons.onetouch.messageencode":
		s := d.Decode(buf, base).(uint8)
		if t.OneTouchAccessEntry.Mode == tytera.SystemType_DIGITAL {
			if t.OneTouchAccessEntry.CallType == tytera.OneTouchCallType_TEXT_MESSAGE {
				t.OneTouchAccessEntry.MessagePreset = uint32(s)
			}
		} else if t.OneTouchAccessEntry.Mode == tytera.SystemType_ANALOG {
			t.OneTouchAccessEntry.DtmfEncodePreset = uint32(s)
		}
	case "com.tytera.buttons.onetouch.call":
		s := d.Decode(buf, base).(uint8)
		if t.OneTouchAccessEntry.Mode == tytera.SystemType_DIGITAL {
			t.OneTouchAccessEntry.ContactIndex = uint32(s)
		}
	}
}
