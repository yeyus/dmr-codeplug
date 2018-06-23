package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
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

func GetZonesGroup() ZonesGroup {
	m := ZonesGroup{
		EntityID: "com.tytera.zones",
		Base:     0x14AE0,
		Length:   0x3E80,
		Zones:    tytera.Zones{},
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x40,
			Records:      250,
			Decoder:      GetZoneEntryDecoder(),
		},
	}

	return m
}

func (t *ZonesGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}

	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.ZoneEntry
	for k, v := range value.([]interface{}) {
		entry := v.(tytera.ZoneEntry)
		arr = append(arr, &entry)
		m[fmt.Sprintf(d.GetEntityID(), k)] = fmt.Sprintf("%+v", entry)
	}

	t.Zones.Entries = arr

	return
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
