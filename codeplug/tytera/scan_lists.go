package tytera

import (
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type ScanListGroup struct {
	EntityID  string
	Base      uint32
	Length    uint32
	Decoders  []encoding.Decoder
	ScanLists tytera.ScanLists
}

func GetScanListGroup() *ScanListGroup {
	m := ScanListGroup{
		EntityID:  "com.tytera.scanLists",
		Base:      0x18960,
		Length:    0x6590,
		ScanLists: tytera.ScanLists{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(tytera.ScanListEntry)

		return !(e.Name == "" &&
			e.PriorityChannel_1 == 0xFFFF &&
			e.PriorityChannel_2 == 0xFFFF &&
			e.TxChannel == 0xFFFF)
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x68,
			Records:      250,
			Decoder:      GetScanListEntryDecoder(),
			Predicate:    predicate,
		},
	}

	return &m
}

func (t *ScanListGroup) GetEntityID() string {
	return t.EntityID
}

func (t *ScanListGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *ScanListGroup) Decode(buf []byte, base uint32) interface{} {
	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.ScanListEntry
	for _, v := range value.([]interface{}) {
		entry := v.(tytera.ScanListEntry)
		arr = append(arr, &entry)
	}

	t.ScanLists.Entries = arr

	return t.ScanLists
}

type ScanListEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.ScanListEntry
}

func GetScanListEntryDecoder() ScanListEntryDecoder {
	m := ScanListEntryDecoder{
		EntityID: "com.tytera.scanLists[%d]",
		Entry:    tytera.ScanListEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.UTF16StringDecoder{
			EntityID:  "name",
			Offset:    0,
			Length:    32,
			Endianess: base.LittleEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "priorityChannel1",
			Offset:    32,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "priorityChannel2",
			Offset:    34,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "txChannel",
			Offset:    36,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "signalHoldTime",
			Offset:   39,
		},
		&base.ByteDecoder{
			EntityID: "prioritySampleTime",
			Offset:   40,
		},
		&encoding.RepeatedDecoder{
			Offset:       42,
			RecordLength: 2,
			Records:      31,
			Decoder: &base.Uint64Decoder{
				EntityID:  "channelIndex[%d]",
				Offset:    0,
				Length:    2,
				Endianess: base.BigEndian,
			},
		},
	}

	return m
}

func (t ScanListEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t ScanListEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t ScanListEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *ScanListEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "name":
		s := d.Decode(buf, base).(string)
		t.Entry.Name = string(s)
	case "priorityChannel1":
		s := d.Decode(buf, base).(uint64)
		t.Entry.PriorityChannel_1 = uint32(s)
	case "priorityChannel2":
		s := d.Decode(buf, base).(uint64)
		t.Entry.PriorityChannel_2 = uint32(s)
	case "txChannel":
		s := d.Decode(buf, base).(uint64)
		t.Entry.TxChannel = uint32(s)
	case "signalHoldTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.SignalHoldTime = uint32(s) * 25
	case "prioritySampleTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.PrioritySampleTime = uint32(s) * 250
	case "channelIndex[%d]":
		s := d.Decode(buf, base).([]interface{})
		var arr []uint32
		for _, v := range s {
			arr = append(arr, uint32(v.(uint64)))
		}

		t.Entry.ChannelIndex = arr
	}
}
