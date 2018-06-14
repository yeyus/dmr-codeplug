package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type EmergencySystemsGroup struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Systems  tytera.EmergencySystems
}

func GetEmergencySystemsGroup() EmergencySystemsGroup {
	m := EmergencySystemsGroup{
		EntityID: "com.tytera.emergency",
		Base:     0x5B70,
		Length:   0x510,
		Systems:  tytera.EmergencySystems{},
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
		},
	}

	return m
}

func (t *EmergencySystemsGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}
	for _, d := range t.Decoders {
		value := d.Decode(buf, base+t.Base)
		switch id := d.GetEntityID(); id {
		case "com.tytera.emergency.entries[%d]":
			var arr []*tytera.EmergencySystemEntry
			for k, v := range value.([]interface{}) {
				entry := v.(tytera.EmergencySystemEntry)
				arr = append(arr, &entry)
				m[fmt.Sprintf(d.GetEntityID(), k)] = fmt.Sprintf("%+v", entry)
			}

			t.Systems.Entries = arr
		default:
			m[id] = fmt.Sprintf("%s", value)
			t.mapValue(d, buf, base+t.Base)
		}

	}

	return
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
