package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type RxGroupListGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Groups   tytera.RxGroups
}

func GetRxGroupListGroup() RxGroupListGroup {
	m := RxGroupListGroup{
		EntityID: "com.tytera.rxGroup",
		Base:     0xED20,
		Length:   0x5DC0,
		Groups:   tytera.RxGroups{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(tytera.RxGroupEntry)

		return e.Name != ""
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x60,
			Records:      250,
			Decoder:      GetRxGroupEntryDecoder(),
			Predicate:    predicate,
		},
	}

	return m
}

func (t *RxGroupListGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}

	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.RxGroupEntry
	for k, v := range value.([]interface{}) {
		entry := v.(tytera.RxGroupEntry)
		arr = append(arr, &entry)
		m[fmt.Sprintf(d.GetEntityID(), k)] = fmt.Sprintf("%+v", entry)
	}

	t.Groups.Entries = arr

	return
}

type RxGroupEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.RxGroupEntry
}

func GetRxGroupEntryDecoder() RxGroupEntryDecoder {
	m := RxGroupEntryDecoder{
		EntityID: "com.tytera.rxGroup[%d]",
		Entry:    tytera.RxGroupEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.UTF16StringDecoder{
			EntityID:  "name",
			Offset:    0,
			Length:    32,
			Endianess: base.LittleEndian,
		},
		&encoding.RepeatedDecoder{
			Offset:       0x20,
			RecordLength: 2,
			Records:      32,
			Decoder: &base.Uint64Decoder{
				EntityID:  "contactIndex[%d]",
				Offset:    0,
				Length:    2,
				Endianess: base.BigEndian,
			},
		},
	}

	return m
}

func (t RxGroupEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t RxGroupEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t RxGroupEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *RxGroupEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "name":
		s := d.Decode(buf, base).(string)
		t.Entry.Name = string(s)
	case "contactIndex[%d]":
		s := d.Decode(buf, base).([]interface{})
		var arr []uint32
		for _, v := range s {
			arr = append(arr, uint32(v.(uint64)))
		}

		t.Entry.ContactIndex = arr
	}
}
