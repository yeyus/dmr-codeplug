package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type ChannelsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Channels tytera.Channels
}

func GetChannelsGroup() *ChannelsGroup {
	m := ChannelsGroup{
		EntityID: "com.tytera.channels",
		Base:     0x1EF00,
		Length:   0xFA00,
		Channels: tytera.Channels{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(tytera.ChannelEntry)

		return e.Name != ""
	}

	m.Decoders = []encoding.Decoder{
		&encoding.RepeatedDecoder{
			Offset:       0,
			RecordLength: 0x40,
			Records:      1000,
			Decoder:      GetChannelEntryDecoder(),
			Predicate:    predicate,
		},
	}

	return &m
}

func (t *ChannelsGroup) GetEntityID() string {
	return t.EntityID
}

func (t *ChannelsGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *ChannelsGroup) Decode(buf []byte, base uint32) interface{} {
	d := t.Decoders[0]
	value := d.Decode(buf, base+t.Base)

	var arr []*tytera.ChannelEntry
	for _, v := range value.([]interface{}) {
		entry := v.(tytera.ChannelEntry)
		arr = append(arr, &entry)
	}

	t.Channels.Entries = arr

	return t.Channels
}

type ChannelEntryDecoder struct {
	EntityID string
	Decoders []encoding.Decoder
	Entry    tytera.ChannelEntry
}

func GetChannelEntryDecoder() ChannelEntryDecoder {
	m := ChannelEntryDecoder{
		EntityID: "com.tytera.channels[%d]",
		Entry:    tytera.ChannelEntry{},
	}

	m.Decoders = []encoding.Decoder{
		&base.BitMaskDecoder{
			EntityID: "mode",
			Offset:   0,
			BitMask:  0x03,
		},
		&base.BitDecoder{
			EntityID:  "bandwidth",
			Offset:    0,
			BitOffset: 3,
		},
		&base.BitDecoder{
			EntityID:  "autoScan",
			Offset:    0,
			BitOffset: 4,
		},
		&base.BitDecoder{
			EntityID:  "squelch",
			Offset:    0,
			BitOffset: 5,
		},
		&base.BitDecoder{
			EntityID:  "loneWorker",
			Offset:    0,
			BitOffset: 7,
		},
		&base.BitMaskDecoder{
			EntityID: "talkaround",
			Offset:   1,
			BitMask:  0x01,
		},
		&base.BitDecoder{
			EntityID:  "rxOnly",
			Offset:    1,
			BitOffset: 1,
		},
		&base.BitMaskDecoder{
			EntityID: "slot",
			Offset:   1,
			BitMask:  0x0C,
		},
		&base.BitMaskDecoder{
			EntityID: "colorCode",
			Offset:   1,
			BitMask:  0xF0,
		},
		&base.BitMaskDecoder{
			EntityID: "keyIndex",
			Offset:   2,
			BitMask:  0x0F,
		},
		&base.BitMaskDecoder{
			EntityID: "privacyType",
			Offset:   2,
			BitMask:  0x30,
		},
		&base.BitDecoder{
			EntityID:  "privateCall",
			Offset:    2,
			BitOffset: 6,
		},
		&base.BitDecoder{
			EntityID:  "dataCall",
			Offset:    2,
			BitOffset: 7,
		},
		&base.BitMaskDecoder{
			EntityID: "rxRefFrequency",
			Offset:   3,
			BitMask:  0x03,
		},
		&base.BitDecoder{
			EntityID:  "emergencyAlarmAck",
			Offset:    3,
			BitOffset: 3,
		},
		&base.BitDecoder{
			EntityID:  "compressedUDPHeader",
			Offset:    3,
			BitOffset: 6,
		},
		&base.BitDecoder{
			EntityID:  "displayPTTId",
			Offset:    3,
			BitOffset: 7,
		},
		&base.BitMaskDecoder{
			EntityID: "txRefFrequency",
			Offset:   4,
			BitMask:  0x03,
		},
		&base.BitDecoder{
			EntityID:  "reverseBurst",
			Offset:    4,
			BitOffset: 2,
		},
		&base.BitDecoder{
			EntityID:  "qtReverse",
			Offset:    4,
			BitOffset: 3,
		},
		&base.BitDecoder{
			EntityID:  "vox",
			Offset:    4,
			BitOffset: 4,
		},
		&base.BitMaskDecoder{
			EntityID: "power",
			Offset:   4,
			BitMask:  0x20,
		},
		&base.BitMaskDecoder{
			EntityID: "admitCriteria",
			Offset:   4,
			BitMask:  0xC0,
		},
		&base.Uint64Decoder{
			EntityID:  "contact",
			Offset:    6,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "tot",
			Offset:   8,
		},
		&base.ByteDecoder{
			EntityID: "totRekey",
			Offset:   9,
		},
		&base.ByteDecoder{
			EntityID: "emergencySystem",
			Offset:   10,
		},
		&base.ByteDecoder{
			EntityID: "scanList",
			Offset:   11,
		},
		&base.ByteDecoder{
			EntityID: "rxGroup",
			Offset:   12,
		},
		&base.BitDecoder{
			EntityID:  "analog1",
			Offset:    14,
			BitOffset: 0,
		},
		&base.BitDecoder{
			EntityID:  "analog2",
			Offset:    14,
			BitOffset: 1,
		},
		&base.BitDecoder{
			EntityID:  "analog3",
			Offset:    14,
			BitOffset: 2,
		},
		&base.BitDecoder{
			EntityID:  "analog4",
			Offset:    14,
			BitOffset: 3,
		},
		&base.BitDecoder{
			EntityID:  "analog5",
			Offset:    14,
			BitOffset: 4,
		},
		&base.BitDecoder{
			EntityID:  "analog6",
			Offset:    14,
			BitOffset: 5,
		},
		&base.BitDecoder{
			EntityID:  "analog7",
			Offset:    14,
			BitOffset: 6,
		},
		&base.BitDecoder{
			EntityID:  "analog8",
			Offset:    14,
			BitOffset: 7,
		},
		&base.BCDDecoder{
			EntityID:  "rxFrequency",
			Offset:    16,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "txFrequency",
			Offset:    20,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.BCDToneDecoder{
			EntityID: "toneDecoding",
			Offset:   24,
		},
		&base.BCDToneDecoder{
			EntityID: "toneEncoding",
			Offset:   26,
		},
		&base.ByteDecoder{
			EntityID: "rxSignaling",
			Offset:   28,
		},
		&base.ByteDecoder{
			EntityID: "txSignaling",
			Offset:   29,
		},
		&base.UTF16StringDecoder{
			EntityID:  "name",
			Offset:    32,
			Length:    32,
			Endianess: base.LittleEndian,
		},
	}

	return m
}

func (t ChannelEntryDecoder) GetEntityID() string {
	return t.EntityID
}

func (t ChannelEntryDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t ChannelEntryDecoder) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.Entry
}

func (t *ChannelEntryDecoder) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "mode":
		s := d.Decode(buf, base).(uint8)
		t.Entry.Mode = tytera.ChannelMode(s)
	case "bandwidth":
		s := d.Decode(buf, base).(bool)
		if s {
			t.Entry.Bandwidth = tytera.Bandwidth_WIDE
		} else {
			t.Entry.Bandwidth = tytera.Bandwidth_NARROW
		}
	case "autoScan":
		s := d.Decode(buf, base).(bool)
		t.Entry.AutoScan = s
	case "squelch":
		s := d.Decode(buf, base).(bool)
		if s {
			t.Entry.Squelch = tytera.Squelch_NORMAL
		} else {
			t.Entry.Squelch = tytera.Squelch_TIGHT
		}
	case "loneWorker":
		s := d.Decode(buf, base).(bool)
		t.Entry.LoneWorker = s
	case "talkaround":
		s := d.Decode(buf, base).(uint8)
		if s == 1 {
			t.Entry.Talkaround = true
		} else {
			t.Entry.Talkaround = false
		}
	case "rxOnly":
		s := d.Decode(buf, base).(bool)
		t.Entry.RxOnly = s
	case "slot":
		s := d.Decode(buf, base).(uint8)
		t.Entry.Slot = tytera.Slot(int(s))
	case "colorCode":
		s := d.Decode(buf, base).(uint8)
		t.Entry.ColorCode = uint32(s & 0x0F)
	case "keyIndex":
		s := d.Decode(buf, base).(uint8)
		t.Entry.KeyIndex = uint32(s)
	case "privacyType":
		s := d.Decode(buf, base).(uint8)
		t.Entry.PrivacyType = tytera.CipherSystem(int(s))
	case "privateCall":
		s := d.Decode(buf, base).(bool)
		t.Entry.PrivateCall = s
	case "dataCall":
		s := d.Decode(buf, base).(bool)
		t.Entry.DataCall = s
	case "rxRefFrequency":
		s := d.Decode(buf, base).(uint8)
		t.Entry.RxRefFrequency = tytera.ReferenceFrequency(int(s))
	case "emergencyAlarmAck":
		s := d.Decode(buf, base).(bool)
		t.Entry.EmergencyAlarmAck = s
	case "compressedUDPHeader":
		s := d.Decode(buf, base).(bool)
		t.Entry.CompressedUdpDataHeader = !s
	case "displayPTTId":
		s := d.Decode(buf, base).(bool)
		t.Entry.DisplayPttId = !s
	case "txRefFrequency":
		s := d.Decode(buf, base).(uint8)
		t.Entry.TxRefFrequency = tytera.ReferenceFrequency(int(s))
	case "reverseBurst":
		s := d.Decode(buf, base).(bool)
		t.Entry.ReverseBurst = s
	case "qtReverse":
		s := d.Decode(buf, base).(bool)
		t.Entry.QtReverse = s
	case "vox":
		s := d.Decode(buf, base).(bool)
		t.Entry.Vox = s
	case "power":
		s := d.Decode(buf, base).(uint8)
		t.Entry.Power = tytera.PowerSetting(int(s))
	case "admitCriteria":
		s := d.Decode(buf, base).(uint8)
		t.Entry.AdmitCriteria = tytera.AdmitCriteria(int(s))
	case "contact":
		s := d.Decode(buf, base).(uint64)
		t.Entry.Contact = uint32(s)
	case "tot":
		s := d.Decode(buf, base).(uint8)
		t.Entry.Tot = uint32(s) * 15
	case "totRekey":
		s := d.Decode(buf, base).(uint8)
		t.Entry.TotRekeyDelay = uint32(s)
	case "emergencySystem":
		s := d.Decode(buf, base).(uint8)
		t.Entry.EmergencySystem = uint32(s)
	case "scanList":
		s := d.Decode(buf, base).(uint8)
		t.Entry.ScanList = uint32(s)
	case "rxGroup":
		s := d.Decode(buf, base).(uint8)
		t.Entry.RxGroup = uint32(s)
	case "analog1":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_1 = s
	case "analog2":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_2 = s
	case "analog3":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_3 = s
	case "analog4":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_4 = s
	case "analog5":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_5 = s
	case "analog6":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_6 = s
	case "analog7":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_7 = s
	case "analog8":
		s := d.Decode(buf, base).(bool)
		t.Entry.AnalogDecode_8 = s
	case "rxFrequency":
		s := d.Decode(buf, base).(uint64)
		t.Entry.RxFrequency = s * 10
	case "txFrequency":
		s := d.Decode(buf, base).(uint64)
		t.Entry.TxFrequency = s * 10
	case "toneDecoding":
		s := d.Decode(buf, base).(tytera.Tone)
		t.Entry.DecodingTone = &s
	case "toneEncoding":
		s := d.Decode(buf, base).(tytera.Tone)
		t.Entry.EncodingTone = &s
	case "rxSignaling":
		s := d.Decode(buf, base).(uint8)
		t.Entry.RxSignaling = tytera.SignalingSystem(int(s))
	case "txSignaling":
		s := d.Decode(buf, base).(uint8)
		t.Entry.TxSignaling = tytera.SignalingSystem(int(s))
	case "name":
		s := d.Decode(buf, base).(string)
		t.Entry.Name = string(s)
	}
}
