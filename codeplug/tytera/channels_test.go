package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"reflect"
	"testing"
)

func TestChannelsParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	cs := GetChannelsGroup()

	b, err := json.MarshalIndent(cs.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", cs.Channels)
}

func TestChannelsProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	cs := GetChannelsGroup()
	cs.Decode(content[:], 0x125)

	if len(cs.Channels.Entries) != 79 {
		t.Errorf("expected number of channels to be 79, but got %d", len(cs.Channels.Entries))
	}

	compareChannel(t, cs, 0, tytera.ChannelEntry{
		Mode:                    tytera.ChannelMode_DIGITAL_CHANNEL,
		Bandwidth:               tytera.Bandwidth_NARROW,
		AutoScan:                false,
		Squelch:                 tytera.Squelch_NORMAL,
		LoneWorker:              false,
		Talkaround:              false,
		RxOnly:                  false,
		Slot:                    tytera.Slot_SLOT1,
		ColorCode:               1,
		PrivacyType:             tytera.CipherSystem_UNSET_CIPHER,
		KeyIndex:                0,
		PrivateCall:             false,
		DataCall:                true,
		RxRefFrequency:          tytera.ReferenceFrequency_LOW_REFERENCE,
		EmergencyAlarmAck:       false,
		CompressedUdpDataHeader: false,
		DisplayPttId:            false,
		TxRefFrequency:          tytera.ReferenceFrequency_LOW_REFERENCE,
		ReverseBurst:            true,
		QtReverse:               false,
		Vox:                     false,
		Power:                   tytera.PowerSetting_LOW_POWER,
		AdmitCriteria:           tytera.AdmitCriteria_ADMIT_ALWAYS,
		Contact:                 2,
		Tot:                     180,
		TotRekeyDelay:           0,
		EmergencySystem:         0,
		ScanList:                0,
		RxGroup:                 0,
		RxFrequency:             437000000,
		TxFrequency:             437000000,
		DecodingTone: &tytera.Tone{
			Type: tytera.ToneSystem_NO_TONE,
		},
		EncodingTone: &tytera.Tone{
			Type: tytera.ToneSystem_NO_TONE,
		},
		TxSignaling: tytera.SignalingSystem_NO_SIGNALING,
		RxSignaling: tytera.SignalingSystem_NO_SIGNALING,
		Name:        "DMR Call",
	})

	compareChannel(t, cs, 78, tytera.ChannelEntry{
		Mode:                    tytera.ChannelMode_ANALOG_CHANNEL,
		Bandwidth:               tytera.Bandwidth_WIDE,
		AutoScan:                true,
		Squelch:                 tytera.Squelch_TIGHT,
		LoneWorker:              true,
		Talkaround:              true,
		RxOnly:                  true,
		Slot:                    tytera.Slot_SLOT1,
		ColorCode:               1,
		PrivacyType:             tytera.CipherSystem_UNSET_CIPHER,
		KeyIndex:                0,
		PrivateCall:             false,
		DataCall:                false,
		RxRefFrequency:          tytera.ReferenceFrequency_MEDIUM_REFERENCE,
		EmergencyAlarmAck:       false,
		CompressedUdpDataHeader: false,
		DisplayPttId:            false,
		TxRefFrequency:          tytera.ReferenceFrequency_MEDIUM_REFERENCE,
		ReverseBurst:            true,
		QtReverse:               false,
		Vox:                     true,
		Power:                   tytera.PowerSetting_HIGH_POWER,
		AdmitCriteria:           tytera.AdmitCriteria_COLOR_CODE,
		Contact:                 0,
		Tot:                     150,
		TotRekeyDelay:           10,
		EmergencySystem:         0,
		ScanList:                2,
		RxGroup:                 0,
		RxFrequency:             420123450,
		TxFrequency:             426123450,
		DecodingTone: &tytera.Tone{
			Type:      tytera.ToneSystem_CTCSS,
			Frequency: 1713,
		},
		EncodingTone: &tytera.Tone{
			Type: tytera.ToneSystem_DCS_NORMAL,
			Code: 155,
		},
		TxSignaling:    tytera.SignalingSystem_DTMF_3,
		RxSignaling:    tytera.SignalingSystem_DTMF_2,
		Name:           "Analog Test",
		AnalogDecode_2: true,
		AnalogDecode_7: true,
	})
}

func compareChannel(t *testing.T, c ChannelsGroup, idx int, x tytera.ChannelEntry) {
	var a tytera.ChannelEntry = *c.Channels.Entries[idx]

	if !reflect.DeepEqual(a, x) {
		t.Errorf("expected channel index %d to be equal to => \n Expected: (type %T)\n %+v\n Got: (type %T)\n %+v", idx, x, x, a, a)
	}
}
