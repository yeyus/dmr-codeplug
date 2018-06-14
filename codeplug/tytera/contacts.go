package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type ContactsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Contacts tytera.Contacts
}

func GetContactsGroup() ContactsGroup {
	m := ContactsGroup{
		EntityID: "com.tytera.contacts",
		Base:     0x6080,
		Length:   0x8CA0,
		Contacts: tytera.Contacts{},
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x24,
			Records:      1000,
			Decoder:      GetContactEntryDecoder(),
		},
	}

	return m
}

func (t *ContactsGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}

	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.ContactEntry
	for k, v := range value.([]interface{}) {
		entry := v.(tytera.ContactEntry)
		arr = append(arr, &entry)
		m[fmt.Sprintf(d.GetEntityID(), k)] = fmt.Sprintf("%+v", entry)
	}

	t.Contacts.Entries = arr

	return
}

type ContactEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.ContactEntry
}

func GetContactEntryDecoder() ContactEntryDecoder {
	m := ContactEntryDecoder{
		EntityID: "com.tytera.contacts[%d]",
		Entry:    tytera.ContactEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.Uint64Decoder{
			EntityID:  "id",
			Offset:    0,
			Length:    3,
			Endianess: base.BigEndian,
		},
		&base.BitMaskDecoder{
			EntityID: "callType",
			Offset:   3,
			BitMask:  0x0F,
		},
		&base.BitDecoder{
			EntityID:  "callReceiveTone",
			Offset:    3,
			BitOffset: 5,
		},
		&base.UTF16StringDecoder{
			EntityID:  "name",
			Offset:    4,
			Length:    32,
			Endianess: base.LittleEndian,
		},
	}

	return m
}

func (t ContactEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t ContactEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t ContactEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *ContactEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "id":
		s := d.Decode(buf, base).(uint64)
		t.Entry.Id = uint32(s)
	case "callType":
		s := d.Decode(buf, base).(uint8)
		t.Entry.CallType = tytera.CallType(uint32(s))
	case "callReceiveTone":
		s := d.Decode(buf, base).(bool)
		t.Entry.CallReceiveTone = s
	case "name":
		s := d.Decode(buf, base).(string)
		t.Entry.Name = string(s)
	}
}
