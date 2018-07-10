package tytera

import (
	"github.com/yeyus/dmr-codeplug/pkg/encoding"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
)

type MD390Codeplug struct {
	EntityID string
	Base     uint32
	Length   uint32
	Decoders []encoding.Decoder
	Codeplug tytera.TyteraCodeplug
}

func GetMD390Codeplug() MD380Codeplug {

	c := MD380Codeplug{
		EntityID: "com.tytera",
		Base:     0,
		Length:   0x40100,
		Codeplug: tytera.TyteraCodeplug{},
	}

	c.Decoders = []encoding.Decoder{
		GetBasicInformationGroup(),
		GetGeneralSettingsGroup(),
		GetMenuItemsGroup(),
		GetButtonsGroup(),
		GetMessagePresetsGroup(),
		GetPrivacySettingsGroup(),
		GetEmergencySystemsGroup(),
		GetContactsGroup(),
		GetRxGroupListGroup(),
		GetZonesGroup(),
		GetScanListGroup(),
		GetChannelsGroup(),
		GetDTMFGroup(),
		GetGPSGroup(),
	}

	return c
}

func (c *MD390Codeplug) Decode(buf []byte, base uint32) interface{} {
	for _, d := range c.Decoders {
		c.mapValue(d, buf, base)
	}

	return c.Codeplug
}

func (c *MD390Codeplug) GetEntityID() string {
	return c.EntityID
}

func (c *MD390Codeplug) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (c *MD390Codeplug) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.basic":
		s := d.Decode(buf, base).(tytera.BasicInformation)
		c.Codeplug.BasicInformation = &s
	case "com.tytera.settings":
		s := d.Decode(buf, base).(tytera.GeneralSettings)
		c.Codeplug.GeneralSettings = &s
	case "com.tytera.menuItems":
		s := d.Decode(buf, base).(tytera.MenuItems)
		c.Codeplug.MenuItems = &s
	case "com.tytera.buttons":
		s := d.Decode(buf, base).(tytera.ButtonDefinitions)
		c.Codeplug.ButtonDefinitions = &s
	case "com.tytera.messages":
		s := d.Decode(buf, base).(tytera.MessagePresets)
		c.Codeplug.MessagePresets = &s
	case "com.tytera.privacy":
		s := d.Decode(buf, base).(tytera.PrivacySettings)
		c.Codeplug.PrivacySettings = &s
	case "com.tytera.emergency":
		s := d.Decode(buf, base).(tytera.EmergencySystems)
		c.Codeplug.EmergencySystems = &s
	case "com.tytera.contacts":
		s := d.Decode(buf, base).(tytera.Contacts)
		c.Codeplug.Contacts = &s
	case "com.tytera.rxGroup":
		s := d.Decode(buf, base).(tytera.RxGroups)
		c.Codeplug.RxGroups = &s
	case "com.tytera.zones":
		s := d.Decode(buf, base).(tytera.Zones)
		c.Codeplug.Zones = &s
	case "com.tytera.scanLists":
		s := d.Decode(buf, base).(tytera.ScanLists)
		c.Codeplug.ScanLists = &s
	case "com.tytera.channels":
		s := d.Decode(buf, base).(tytera.Channels)
		c.Codeplug.Channels = &s
	case "com.tytera.dtmf":
		s := d.Decode(buf, base).(tytera.DTMFSettings)
		c.Codeplug.Dtmf = &s
	case "com.tytera.gps":
		s := d.Decode(buf, base).(tytera.GPSSystems)
		c.Codeplug.Gps = &s

	}
}
