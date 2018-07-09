package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/pkg/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type GeneralSettingsGroup struct {
	EntityID        string
	Base            uint32
	Length          uint32
	Decoders        []encoding.Decoder
	GeneralSettings tytera.GeneralSettings
}

func GetGeneralSettingsGroup() *GeneralSettingsGroup {
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
			EntityID:  "com.tytera.settings.disableLeds",
			Offset:    64,
			BitOffset: 2,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.monitorType",
			Offset:    64,
			BitOffset: 4,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.batSavePreamble",
			Offset:    65,
			BitOffset: 0,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.batSaveReceiveMode",
			Offset:    65,
			BitOffset: 1,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disableTones",
			Offset:    65,
			BitOffset: 2,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disableChFreeTone",
			Offset:    65,
			BitOffset: 4,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.disablePasswordAndLock",
			Offset:    65,
			BitOffset: 5,
		},
		&base.BitMaskDecoder{
			EntityID: "com.tytera.settings.talkPermitTone",
			Offset:   65,
			BitMask:  0xC0,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.settings.introScreenMode",
			Offset:    66,
			BitOffset: 4,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.settings.radioId",
			Offset:    68,
			Length:    3,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.txPreamble",
			Offset:   72,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.groupCallHang",
			Offset:   73,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.privateCallHang",
			Offset:   74,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.voxSensitivity",
			Offset:   75,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.rxLowBattery",
			Offset:   78,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.callAlertTone",
			Offset:   79,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.loneWorkerResponse",
			Offset:   80,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.loneWorkerReminder",
			Offset:   81,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.scanDigitalHangTime",
			Offset:   83,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.scanAnalogHangTime",
			Offset:   84,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.backlightTimeout",
			Offset:   85,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.keypadLock",
			Offset:   86,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.settings.operationMode",
			Offset:   87,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.settings.powerOnPassword",
			Offset:    88,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.settings.radioProgrammingPassword",
			Offset:    92,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.ASCIIStringDecoder{
			EntityID: "com.tytera.settings.pcProgrammingPassword",
			Offset:   96,
			Length:   8,
		},
		&base.UTF16StringDecoder{
			EntityID:  "com.tytera.settings.radioName",
			Offset:    112,
			Length:    32,
			Endianess: base.LittleEndian,
		},
	}

	return &g
}

func (t *GeneralSettingsGroup) GetEntityID() string {
	return t.EntityID
}

func (t *GeneralSettingsGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *GeneralSettingsGroup) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base+t.Base)
	}

	return t.GeneralSettings
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
	case "com.tytera.settings.backlightTimeout":
		s := d.Decode(buf, base).(uint8)
		t.GeneralSettings.BacklightTimeout = uint32(s) * 5
	case "com.tytera.settings.keypadLock":
		s := d.Decode(buf, base).(uint8)
		// 0xFF is manual
		if s == 0xFF {
			t.GeneralSettings.KeypadLock = 255
		} else {
			t.GeneralSettings.KeypadLock = uint32(s) * 5
		}
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
		if s == "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" {
			t.GeneralSettings.PcProgrammingPassword = ""
		} else {
			t.GeneralSettings.PcProgrammingPassword = s
		}
	case "com.tytera.settings.radioName":
		s := d.Decode(buf, base).(string)
		t.GeneralSettings.RadioName = s
	}
}
