package tytera

import (
	"github.com/yeyus/dmr-codeplug/proto/tytera"
)

func Parse(contents []byte) (tytera.TyteraCodeplug, error) {

	codeplug := tytera.TyteraCodeplug{}

	// Basic Information
	bi := GetBasicInformationGroup()
	bi.Decode(contents, 0)
	codeplug.BasicInformation = &bi.BasicInformation

	// General Settings
	gs := GetGeneralSettingsGroup()
	gs.Decode(contents, 0)
	codeplug.GeneralSettings = &gs.GeneralSettings

	// Menu Items
	mi := GetMenuItemsGroup()
	mi.Decode(contents, 0)
	codeplug.MenuItems = &mi.MenuItems

	// Button Definitions
	bd := GetButtonsGroup()
	bd.Decode(contents, 0)
	codeplug.ButtonDefinitions = &bd.Buttons

	// Message Presets
	mp := GetMessagePresetsGroup()
	mp.Decode(contents, 0)
	codeplug.MessagePresets = &mp.Messages

	// Privacy Settings
	ps := GetPrivacySettingsGroup()
	ps.Decode(contents, 0)
	codeplug.PrivacySettings = &ps.Privacy

	// Emergency Systems
	es := GetEmergencySystemsGroup()
	es.Decode(contents, 0)
	codeplug.EmergencySystems = &es.Systems

	// Contacts
	cs := GetContactsGroup()
	cs.Decode(contents, 0)
	codeplug.Contacts = &cs.Contacts

	return codeplug, nil
}
