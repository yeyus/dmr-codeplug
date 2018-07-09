package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type EmergencySystemsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Systems  tytera.EmergencySystems
}

func GetEmergencySystemsGroup() *EmergencySystemsGroup {
	m := EmergencySystemsGroup{
		EntityID: "com.tytera.emergency",
		Base:     0x5B70,
		Length:   0x510,
		Systems:  tytera.EmergencySystems{},
	}

	predicate := func(entry interface{}) bool {
		e := entry.(tytera.EmergencySystemEntry)

		return e.SystemName != ""
	}

	m.Decoders = []encoding.Decoder{
		&base.BitDecoder{
			EntityID:  "com.tytera.emergency.radioDisableDecode",
			Offset:    0,
			BitOffset: 0,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.emergency.remoteMonitorDecode",
			Offset:    0,
			BitOffset: 1,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.emergency.emergencyRemoteMonitorDecode",
			Offset:    0,
			BitOffset: 2,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.emergency.remoteMonitorDuration",
			Offset:   1,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.emergency.txSyncWakeup",
			Offset:   2,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.emergency.txWakeupMessageLimit",
			Offset:   3,
		},
		&encoding.RepeatedDecoder{
			Offset:       16,
			RecordLength: 0x28,
			Records:      32,
			Decoder:      GetEmergencySystemEntry(),
			Predicate:    predicate,
		},
	}

	return &m
}

func (t *EmergencySystemsGroup) GetEntityID() string {
	return t.EntityID
}

func (t *EmergencySystemsGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *EmergencySystemsGroup) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		value := d.Decode(buf, base+t.Base)
		switch id := d.GetEntityID(); id {
		case "com.tytera.emergency.entries[%d]":
			var arr []*tytera.EmergencySystemEntry
			for _, v := range value.([]interface{}) {
				entry := v.(tytera.EmergencySystemEntry)
				arr = append(arr, &entry)
			}

			t.Systems.Entries = arr
		default:
			t.mapValue(d, buf, base+t.Base)
		}

	}

	return t.Systems
}

func (t *EmergencySystemsGroup) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.emergency.radioDisableDecode":
		s := d.Decode(buf, base).(bool)
		t.Systems.RadioDisableDecode = s
	case "com.tytera.emergency.remoteMonitorDecode":
		s := d.Decode(buf, base).(bool)
		t.Systems.RemoteMonitorDecode = s
	case "com.tytera.emergency.emergencyRemoteMonitorDecode":
		s := d.Decode(buf, base).(bool)
		t.Systems.EmergencyRemoteMonitorDecode = s
	case "com.tytera.emergency.remoteMonitorDuration":
		s := d.Decode(buf, base).(uint8)
		t.Systems.RemoteMonitorDuration = uint32(s)
	case "com.tytera.emergency.txSyncWakeup":
		s := d.Decode(buf, base).(uint8)
		t.Systems.TxSyncWakeup = uint32(s) * 25
	case "com.tytera.emergency.txWakeupMessageLimit":
		s := d.Decode(buf, base).(uint8)
		t.Systems.TxWakeupMessageLimit = uint32(s)
	}
}
