package tytera

import (
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type MenuItemsGroup struct {
	EntityID  string
	Base      uint32
	Length    uint32
	Decoders  []encoding.Decoder
	MenuItems tytera.MenuItems
}

func GetMenuItemsGroup() *MenuItemsGroup {
	m := MenuItemsGroup{
		EntityID: "com.tytera.menuItems",
		Base:     0x21F0,
		Length:   5,
		MenuItems: tytera.MenuItems{
			Contacts:  &tytera.ContactsMenu{},
			CallLog:   &tytera.CallLogMenu{},
			Scan:      &tytera.ScanMenu{},
			Utilities: &tytera.UtilitiesMenu{},
		},
	}

	m.Decoders = []encoding.Decoder{
		&base.ByteDecoder{
			EntityID: "com.tytera.menuItems.hangTime",
			Offset:   0,
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.radioDisable",
			Offset:    8 / 8,
			BitOffset: 7 - (8 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.radioEnable",
			Offset:    9 / 8,
			BitOffset: 7 - (9 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.remoteMonitor",
			Offset:    10 / 8,
			BitOffset: 7 - (10 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.radioCheck",
			Offset:    11 / 8,
			BitOffset: 7 - (11 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.manualDial",
			Offset:    12 / 8,
			BitOffset: 7 - (12 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.edit",
			Offset:    13 / 8,
			BitOffset: 7 - (13 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.callAlert",
			Offset:    14 / 8,
			BitOffset: 7 - (14 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.textMessages",
			Offset:    15 / 8,
			BitOffset: 7 - (15 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.toneOrAlert",
			Offset:    16 / 8,
			BitOffset: 7 - (16 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.talkaround",
			Offset:    17 / 8,
			BitOffset: 7 - (17 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.callLog.outgoingRadio",
			Offset:    18 / 8,
			BitOffset: 7 - (18 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.callLog.answered",
			Offset:    19 / 8,
			BitOffset: 7 - (19 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.callLog.missed",
			Offset:    20 / 8,
			BitOffset: 7 - (20 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.scan.editList",
			Offset:    21 / 8,
			BitOffset: 7 - (21 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.scan.scan",
			Offset:    22 / 8,
			BitOffset: 7 - (22 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.contacts.programKey",
			Offset:    23 / 8,
			BitOffset: 7 - (23 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.vox",
			Offset:    24 / 8,
			BitOffset: 7 - (24 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.squelch",
			Offset:    26 / 8,
			BitOffset: 7 - (26 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.ledIndicator",
			Offset:    27 / 8,
			BitOffset: 7 - (27 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.keyboardLock",
			Offset:    28 / 8,
			BitOffset: 7 - (28 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.introScreen",
			Offset:    29 / 8,
			BitOffset: 7 - (29 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.backlight",
			Offset:    30 / 8,
			BitOffset: 7 - (30 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.power",
			Offset:    31 / 8,
			BitOffset: 7 - (31 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.gps",
			Offset:    36 / 8,
			BitOffset: 7 - (36 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.programRadio",
			Offset:    37 / 8,
			BitOffset: 7 - (37 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.displayMode",
			Offset:    38 / 8,
			BitOffset: 7 - (38 % 8),
		},
		&base.BitDecoder{
			EntityID:  "com.tytera.menuItems.utilities.passwordAndLock",
			Offset:    39 / 8,
			BitOffset: 7 - (39 % 8),
		},
	}

	return &m
}

func (t *MenuItemsGroup) GetEntityID() string {
	return t.EntityID
}

func (t *MenuItemsGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *MenuItemsGroup) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base+t.Base)
	}

	return t.MenuItems
}

func (t *MenuItemsGroup) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.menuItems.hangTime":
		o := d.Decode(buf, base).(uint8)
		t.MenuItems.HangTime = uint32(o)
	case "com.tytera.menuItems.contacts.radioDisable":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.RadioDisable = o
	case "com.tytera.menuItems.contacts.radioEnable":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.RadioEnable = o
	case "com.tytera.menuItems.contacts.remoteMonitor":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.RemoteMonitor = o
	case "com.tytera.menuItems.contacts.radioCheck":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.RadioCheck = o
	case "com.tytera.menuItems.contacts.manualDial":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.ManualDial = o
	case "com.tytera.menuItems.contacts.edit":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.Edit = o
	case "com.tytera.menuItems.contacts.callAlert":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.CallAlert = o
	case "com.tytera.menuItems.textMessages":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.TextMessages = o
	case "com.tytera.menuItems.utilities.toneOrAlert":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.ToneOrAlert = o
	case "com.tytera.menuItems.utilities.talkaround":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Talkaround = o
	case "com.tytera.menuItems.callLog.outgoingRadio":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.CallLog.OutgoingRadio = o
	case "com.tytera.menuItems.callLog.answered":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.CallLog.Answered = o
	case "com.tytera.menuItems.callLog.missed":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.CallLog.Missed = o
	case "com.tytera.menuItems.scan.editList":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Scan.EditList = o
	case "com.tytera.menuItems.scan.scan":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Scan.Scan = o
	case "com.tytera.menuItems.contacts.programKey":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Contacts.ProgramKey = o
	case "com.tytera.menuItems.utilities.vox":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Vox = o
	case "com.tytera.menuItems.utilities.squelch":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Squelch = o
	case "com.tytera.menuItems.utilities.ledIndicator":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.LedIndicator = o
	case "com.tytera.menuItems.utilities.keyboardLock":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.KeyboardLock = o
	case "com.tytera.menuItems.utilities.introScreen":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.IntroScreen = o
	case "com.tytera.menuItems.utilities.backlight":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Backlight = o
	case "com.tytera.menuItems.utilities.power":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Power = o
	case "com.tytera.menuItems.utilities.gps":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.Gps = !o
	case "com.tytera.menuItems.utilities.programRadio":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.ProgramRadio = !o
	case "com.tytera.menuItems.utilities.displayMode":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.DisplayMode = o
	case "com.tytera.menuItems.utilities.passwordAndLock":
		o := d.Decode(buf, base).(bool)
		t.MenuItems.Utilities.PasswordAndLock = o
	}
}
