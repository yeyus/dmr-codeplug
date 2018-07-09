package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type DTMFDecodeEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.DTMFDecodeEntry
}

func GetDTMFDecodeEntryDecoder() DTMFDecodeEntryDecoder {
	d := DTMFDecodeEntryDecoder{
		EntityID: "com.tytera.dtmf.decode[%d]",
		Entry:    tytera.DTMFDecodeEntry{},
	}

	d.Decoders = []encoding.Decoder{
		&base.DTMFStringDecoder{
			EntityID: "dtmfId",
			Offset:   0,
		},
		&base.ByteDecoder{
			EntityID: "responseType",
			Offset:   16,
		},
		&base.ByteDecoder{
			EntityID: "decodeType",
			Offset:   17,
		},
		&base.ByteDecoder{
			EntityID: "ackEncodeIndex",
			Offset:   18,
		},
		&base.ByteDecoder{
			EntityID: "ackDelayTime",
			Offset:   19,
		},
		&base.Uint64Decoder{
			EntityID:  "revertChannel",
			Offset:    20,
			Length:    2,
			Endianess: base.BigEndian,
		},
	}

	return d
}

func (t DTMFDecodeEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t DTMFDecodeEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t DTMFDecodeEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *DTMFDecodeEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "dtmfId":
		s := d.Decode(buf, base).(string)
		t.Entry.DtmfId = s
	case "responseType":
		s := d.Decode(buf, base).(uint8)
		t.Entry.ResponseType = tytera.ResponseType(int(s))
	case "decodeType":
		s := d.Decode(buf, base).(uint8)
		t.Entry.DecodeType = tytera.DecodeType(int(s))
	case "ackEncodeIndex":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AckEncodeIndex = uint32(s)
	case "ackDelayTime":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AckDelayTime = uint32(s) * 50
	case "revertChannel":
		s := d.Decode(buf, base).(uint64)
		t.Entry.RevertChannelIndex = uint32(s)
	}
}
