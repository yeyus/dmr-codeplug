package tytera

import (
	"fmt"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

type GeneralSettingsGroup struct {
	EntityID        string
	Base            uint32
	Length          uint32
	Decoders        []encoding.Decoder
	GeneralSettings tytera.GeneralSettings
}

func GetGeneralSettingsGroup() GeneralSettingsGroup {
	g := GeneralSettingsGroup{
		EntityID:        "com.tytera.settings",
		Base:            0x2140,
		Length:          0x90,
		GeneralSettings: tytera.GeneralSettings{},
	}

	g.Decoders = []encoding.Decoder{
		&base.UTF16StringDecoder{
			EntityID:  "com.tytera.settings.screenLine[0]",
			Offset:    0,
			Length:    20,
			Endianess: base.LittleEndian,
		},
		&base.UTF16StringDecoder{
			EntityID:  "com.tytera.settings.screenLine[1]",
			Offset:    20,
			Length:    20,
			Endianess: base.LittleEndian,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.monitorType",
			Offset:    515 / 8,
			BitOffset: 7 - (515 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disableLeds",
			Offset:    517 / 8,
			BitOffset: 7 - (517 % 8),
		},
		&base.BitMaskDecoder{
			EntityID: "com.tytera.settings.talkPermitTone",
			Offset:   520 / 8,
			BitMask:  0xC0,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disablePasswordAndLock",
			Offset:    522 / 8,
			BitOffset: 7 - (522 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disableChFreeTone",
			Offset:    523 / 8,
			BitOffset: 7 - (523 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disableTones",
			Offset:    525 / 8,
			BitOffset: 7 - (525 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.batSaveReceiveMode",
			Offset:    526 / 8,
			BitOffset: 7 - (526 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.batSavePreamble",
			Offset:    527 / 8,
			BitOffset: 7 - (527 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.introScreenMode",
			Offset:    531 / 8,
			BitOffset: 7 - (531 % 8),
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.settings.radioId",
			Offset:    544 / 8,
			Length:    3,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.txPreamble",
			Offset:   576 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.groupCallHang",
			Offset:   584 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.privateCallHang",
			Offset:   592 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.voxSensitivity",
			Offset:   600 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.rxLowBattery",
			Offset:   624 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.callAlertTone",
			Offset:   632 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.loneWorkerResponse",
			Offset:   640 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.loneWorkerReminder",
			Offset:   648 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.scanDigitalHangTime",
			Offset:   664 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.scanAnalogHangTime",
			Offset:   672 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.keypadLock",
			Offset:   688 / 8,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.operationMode",
			Offset:   696 / 8,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.settings.powerOnPassword",
			Offset:    704 / 8,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.settings.radioProgrammingPassword",
			Offset:    736 / 8,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.ASCIIStringDecoder{
			EntityID: "com.tytera.settings.pcProgrammingPassword",
			Offset:   704 / 8,
			Length:   4,
		},
		&base.UTF16StringDecoder{
			EntityID:  "com.tytera.settings.radioName",
			Offset:    112,
			Length:    32,
			Endianess: base.LittleEndian,
		},
	}

	return g
}

func (t *GeneralSettingsGroup) Decode(buf []byte, base uint32) (m map[string]string) {
	m = map[string]string{}
	for _, d := range t.Decoders {
		m[d.GetEntityID()] = fmt.Sprintf("%s", d.Decode(buf, base+t.Base))
		t.mapValue(d, buf, base+t.Base)
	}

	return
}

func (t *GeneralSettingsGroup) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.settings.screenLine[0]", "com.tytera.settings.screenLine[1]":
		s := d.Decode(buf, base).(string)
		t.GeneralSettings.ScreenLines = append(t.GeneralSettings.ScreenLines, s)
	case "com.tytera.settings.monitorType":
		s := d.Decode(buf, base).(bool)
		if s {
			t.GeneralSettings.MonitorType = tytera.MonitorType_OPEN_SQUELCH
		} else {
			t.GeneralSettings.MonitorType = tytera.MonitorType_SILENT
		}
	case "com.tytera.settings.disableLeds":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.DisableLeds = !s
	case "com.tytera.settings.talkPermitTone":

		switch s := d.Decode(buf, base).(uint8); s {
		case 0:
			t.GeneralSettings.TalkPermitTone = tytera.SystemType_NONE
		case 1:
			t.GeneralSettings.TalkPermitTone = tytera.SystemType_DIGITAL
		case 2:
			t.GeneralSettings.TalkPermitTone = tytera.SystemType_ANALOG
		case 3:
			t.GeneralSettings.TalkPermitTone = tytera.SystemType_BOTH
		}
	case "com.tytera.settings.disablePasswordAndLock":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.DisablePasswordAndLock = s
	case "com.tytera.settings.disableChFreeTone":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.DisableChFreeTone = s
	case "com.tytera.settings.disableTones":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.DisableTones = !s
	case "com.tytera.settings.batSaveReceiveMode":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.BatSaveReceiveMode = s
	case "com.tytera.settings.batSavePreamble":
		s := d.Decode(buf, base).(bool)
		t.GeneralSettings.BatSavePreamble = s
	case "com.tytera.settings.introScreenMode":
		s := d.Decode(buf, base).(bool)
		if s {
			t.GeneralSettings.IntroScreenMode = tytera.IntroScreenMode_PICTURE
		} else {
			t.GeneralSettings.IntroScreenMode = tytera.IntroScreenMode_CHAR_STRING
		}
	case "com.tytera.settings.radioId":
		s := d.Decode(buf, base).(uint64)
		t.GeneralSettings.RadioId = uint32(s)
	case "com.tytera.settings.txPreamble":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.TxPreamble = uint32(s) * 60
	case "com.tytera.settings.groupCallHang":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.GroupCallHang = uint32(s) * 100
	case "com.tytera.settings.privateCallHang":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.PrivateCallHang = uint32(s) * 100
	case "com.tytera.settings.voxSensitivity":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.VoxSensitivity = uint32(s)
	case "com.tytera.settings.rxLowBattery":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.RxLowBattery = uint32(s) * 5
	case "com.tytera.settings.callAlertTone":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.CallAlertTone = uint32(s)
	case "com.tytera.settings.loneWorkerResponse":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.LoneWorkerResponse = uint32(s)
	case "com.tytera.settings.loneWorkerReminder":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.LoneWorkerReminder = uint32(s)
	case "com.tytera.settings.scanDigitalHangTime":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.ScanDigitalHangTime = uint32(s) * 100
	case "com.tytera.settings.scanAnalogHangTime":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.ScanAnalogHangTime = uint32(s) * 100
	case "com.tytera.settings.keypadLock":
		s := d.Decode(buf, base).(uint8)
		// 0xFF is manual
		t.GeneralSettings.KeypadLock = uint32(s) * 5
	case "com.tytera.settings.operationMode":
		s := d.Decode(buf, base).(uint8)
		if s == 0 {
			t.GeneralSettings.OperationMode = tytera.OperationMode_MR
		} else {
			t.GeneralSettings.OperationMode = tytera.OperationMode_CHANNEL
		}
	case "com.tytera.settings.powerOnPassword":
		s := d.Decode(buf, base).(uint64)
		t.GeneralSettings.PowerOnPassword = uint32(s)
	case "com.tytera.settings.radioProgrammingPassword":
		s := d.Decode(buf, base).(uint64)
		t.GeneralSettings.RadioProgrammingPassword = uint32(s)
	case "com.tytera.settings.pcProgrammingPassword":
		s := d.Decode(buf, base).(string)
		t.GeneralSettings.PcProgrammingPassword = s
	case "com.tytera.settings.radioName":
		s := d.Decode(buf, base).(string)
		t.GeneralSettings.RadioName = s
	}
}