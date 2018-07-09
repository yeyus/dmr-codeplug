package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type ZonesGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Zones    tytera.Zones
}

func GetZonesGroup() *ZonesGroup {
	m := ZonesGroup{
		EntityID: "com.tytera.zones",
		Base:     0x14AE0,
		Length:   0x3E80,
		Zones:    tytera.Zones{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(tytera.ZoneEntry)

		return e.Name != ""
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x40,
			Records:      250,
			Decoder:      GetZoneEntryDecoder(),
			Predicate:    predicate,
		},
	}

	return &m
}

func (t *ZonesGroup) GetEntityID() string {
	return t.EntityID
}

func (t *ZonesGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *ZonesGroup) Decode(buf []byte, base uint32) interface{} {
	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.ZoneEntry
	for _, v := range value.([]interface{}) {
		entry := v.(tytera.ZoneEntry)
		arr = append(arr, &entry)
	}

	t.Zones.Entries = arr

	return t.Zones
}

type ZoneEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.ZoneEntry
}

func GetZoneEntryDecoder() ZoneEntryDecoder {
	m := ZoneEntryDecoder{
		EntityID: "com.tytera.zones[%d]",
		Entry:    tytera.ZoneEntry{},
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
			Records:      16,
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

func (t ZoneEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t ZoneEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t ZoneEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *ZoneEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "name":
		s := d.Decode(buf, base).(string)
		t.Entry.Name = string(s)
	case "channelIndex[%d]":
		s := d.Decode(buf, base).([]interface{})
		var arr []uint32
		for _, v := range s {
			arr = append(arr, uint32(v.(uint64)))
		}

		t.Entry.ChannelIndex = arr
	}
}
