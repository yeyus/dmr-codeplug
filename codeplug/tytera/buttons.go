package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type ButtonsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Buttons  tytera.ButtonDefinitions
}

func GetButtonsGroup() ButtonsGroup {
	m := ButtonsGroup{
		EntityID: "com.tytera.buttons",
		Base:     0x2202,
		Length:   0x3D,
		Buttons:  tytera.ButtonDefinitions{},
	}

	m.Decoders = []encoding.Decoder{
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.longPressDuration",
			Offset:   0xF,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.sideShort1",
			Offset:   0x0,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.sideLong1",
			Offset:   0x1,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.sideShort2",
			Offset:   0x2,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.buttons.sideLong2",
			Offset:   0x3,
		},
		&encoding.RepeatedDecoder{
			Offset:       0x12,
			RecordLength: 4,
			Records:      6,
			Decoder:      GetOneTouchAccessEntry(),
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey0",
			Offset:    0x2A,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey1",
			Offset:    0x2C,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey2",
			Offset:    0x2E,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey3",
			Offset:    0x30,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey4",
			Offset:    0x32,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey5",
			Offset:    0x34,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey6",
			Offset:    0x36,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey7",
			Offset:    0x38,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey8",
			Offset:    0x3A,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.buttons.contactKey9",
			Offset:    0x3C,
			Length:    2,
			Endianess: base.BigEndian,
		},
	}

	return m
}

func (t *ButtonsGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}
	for _, d := range t.Decoders {
		m[d.GetEntityID()] = fmt.Sprintf("%s", d.Decode(buf, base+t.Base))
		t.mapValue(d, buf, base+t.Base)
	}

	return
}

func (t *ButtonsGroup) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.buttons.longPressDuration":
		s := d.Decode(buf, base).(uint8)
		t.Buttons.LongPressDuration = uint32(s) * 250
	case "com.tytera.buttons.sideShort1":
		s := d.Decode(buf, base).(uint8)
		t.Buttons.SideShort_1 = decodeButtonFunction(s)
	case "com.tytera.buttons.sideLong1":
		s := d.Decode(buf, base).(uint8)
		t.Buttons.SideLong_1 = decodeButtonFunction(s)
	case "com.tytera.buttons.sideShort2":
		s := d.Decode(buf, base).(uint8)
		t.Buttons.SideShort_2 = decodeButtonFunction(s)
	case "com.tytera.buttons.sideLong2":
		s := d.Decode(buf, base).(uint8)
		t.Buttons.SideLong_2 = decodeButtonFunction(s)
	case "com.tytera.buttons.onetouch[%d]":
		s := d.Decode(buf, base)
		var arr []*tytera.OneTouchAccessEntry
		for _, ota := range s.([]interface{}) {
			p := ota.(tytera.OneTouchAccessEntry)
			arr = append(arr, &p)
		}

		t.Buttons.OneTouchAccess = arr
	case "com.tytera.buttons.contactKey0":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_0 = uint32(s)
	case "com.tytera.buttons.contactKey1":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_1 = uint32(s)
	case "com.tytera.buttons.contactKey2":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_2 = uint32(s)
	case "com.tytera.buttons.contactKey3":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_3 = uint32(s)
	case "com.tytera.buttons.contactKey4":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_4 = uint32(s)
	case "com.tytera.buttons.contactKey5":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_5 = uint32(s)
	case "com.tytera.buttons.contactKey6":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_6 = uint32(s)
	case "com.tytera.buttons.contactKey7":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_7 = uint32(s)
	case "com.tytera.buttons.contactKey8":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_8 = uint32(s)
	case "com.tytera.buttons.contactKey9":
		s := d.Decode(buf, base).(uint64)
		t.Buttons.ContactKey_9 = uint32(s)
	}
}

func decodeButtonFunction(value uint8) tytera.ButtonFunction {
	v := int32(value)
	return tytera.ButtonFunction(v)
}
