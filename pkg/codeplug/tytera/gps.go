package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type GPSGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	GPS      tytera.GPSSystems
}

func GetGPSGroup() *GPSGroup {
	m := GPSGroup{
		EntityID: "com.tytera.gps",
		Base:     0x3ED40,
		Length:   0x160,
		GPS:      tytera.GPSSystems{},
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x10,
			Records:      16,
			Decoder:      GetGPSEntryDecoder(),
		},
	}

	return &m
}

func (t *GPSGroup) GetEntityID() string {
	return t.EntityID
}

func (t *GPSGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *GPSGroup) Decode(buf []byte, base uint32) interface{} {
	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.GPSSystemEntry
	for _, v := range value.([]interface{}) {
		entry := v.(tytera.GPSSystemEntry)
		arr = append(arr, &entry)
	}

	t.GPS.Entries = arr

	return t.GPS
}

type GPSEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.GPSSystemEntry
}

func GetGPSEntryDecoder() *GPSEntryDecoder {
	m := GPSEntryDecoder{
		EntityID: "com.tytera.gps[%d]",
		Entry:    tytera.GPSSystemEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.Uint64Decoder{
			EntityID:  "revertChannel",
			Offset:    0,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "reportInterval",
			Offset:   2,
		},
		&base.Uint64Decoder{
			EntityID:  "destinationId",
			Offset:    4,
			Length:    2,
			Endianess: base.BigEndian,
		},
	}

	return &m
}

func (t *GPSEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t *GPSEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *GPSEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *GPSEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "revertChannel":
		s := d.Decode(buf, base).(uint64)
		t.Entry.RevertChannelIndex = uint32(s)
	case "reportInterval":
		s := d.Decode(buf, base).(uint8)
		t.Entry.ReportInterval = uint32(s) * 30
	case "destinationId":
		s := d.Decode(buf, base).(uint64)
		t.Entry.DestinationContactIndex = uint32(s)
	}
}
