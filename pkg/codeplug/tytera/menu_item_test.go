package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMenuItemsParsing(t *testing.T) {
	content := getRDTBytes("../../../test/04_Oct_2017.rdt")

	mi := GetMenuItemsGroup()

	m, err := json.MarshalIndent(mi.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(m))

	fmt.Printf("%+v", mi.MenuItems)
}

func TestMenuItemsProto(t *testing.T) {
	content := getRDTBytes("../../../test/04_Oct_2017.rdt")

	mi := GetMenuItemsGroup()

	mi.Decode(content[:], 0x125)

	if mi.MenuItems.HangTime != 10 {
		t.Errorf("expected menu hang time to be 10, got %d", mi.MenuItems.HangTime)
	}

	if mi.MenuItems.TextMessages != true {
		t.Errorf("expected menu text messages to be true, got %v", mi.MenuItems.TextMessages)
	}

	if mi.MenuItems.Contacts.CallAlert != true {
		t.Errorf("expected menu call alert to be true, got %v", mi.MenuItems.Contacts.CallAlert)
	}

	if mi.MenuItems.Contacts.Edit != true {
		t.Errorf("expected menu edit to be true, got %v", mi.MenuItems.Contacts.Edit)
	}

	if mi.MenuItems.Contacts.ManualDial != true {
		t.Errorf("expected menu manual dial to be true, got %v", mi.MenuItems.Contacts.ManualDial)
	}

	if mi.MenuItems.Contacts.RadioCheck != false {
		t.Errorf("expected menu radio check to be false, got %v", mi.MenuItems.Contacts.RadioCheck)
	}

	if mi.MenuItems.Contacts.RemoteMonitor != false {
		t.Errorf("expected menu remote monitor to be false, got %v", mi.MenuItems.Contacts.RemoteMonitor)
	}

	if mi.MenuItems.Contacts.ProgramKey != true {
		t.Errorf("expected menu program key to be true, got %v", mi.MenuItems.Contacts.ProgramKey)
	}

	if mi.MenuItems.Contacts.RadioEnable != false {
		t.Errorf("expected menu radio enable to be false, got %v", mi.MenuItems.Contacts.RadioEnable)
	}

	if mi.MenuItems.Contacts.RadioDisable != false {
		t.Errorf("expected menu radio disable to be false, got %v", mi.MenuItems.Contacts.RadioDisable)
	}

	if mi.MenuItems.CallLog.Missed != true {
		t.Errorf("expected menu missed to be true, got %v", mi.MenuItems.CallLog.Missed)
	}

	if mi.MenuItems.CallLog.Answered != true {
		t.Errorf("expected menu answered to be true, got %v", mi.MenuItems.CallLog.Answered)
	}

	if mi.MenuItems.CallLog.OutgoingRadio != true {
		t.Errorf("expected menu outgoing radio to be true, got %v", mi.MenuItems.CallLog.OutgoingRadio)
	}

	if mi.MenuItems.Utilities.Talkaround != true {
		t.Errorf("expected menu talkaround to be true, got %v", mi.MenuItems.Utilities.Talkaround)
	}

	if mi.MenuItems.Utilities.ToneOrAlert != true {
		t.Errorf("expected menu tone or alert to be true, got %v", mi.MenuItems.Utilities.ToneOrAlert)
	}

	if mi.MenuItems.Utilities.Power != true {
		t.Errorf("expected menu power to be true, got %v", mi.MenuItems.Utilities.Power)
	}

	if mi.MenuItems.Utilities.Backlight != true {
		t.Errorf("expected menu backlight to be true, got %v", mi.MenuItems.Utilities.Backlight)
	}

	if mi.MenuItems.Utilities.IntroScreen != true {
		t.Errorf("expected menu intro screen to be true, got %v", mi.MenuItems.Utilities.IntroScreen)
	}

	if mi.MenuItems.Utilities.KeyboardLock != true {
		t.Errorf("expected menu keyboard lock to be true, got %v", mi.MenuItems.Utilities.KeyboardLock)
	}

	if mi.MenuItems.Utilities.LedIndicator != true {
		t.Errorf("expected menu led indicator to be true, got %v", mi.MenuItems.Utilities.LedIndicator)
	}

	if mi.MenuItems.Utilities.Squelch != true {
		t.Errorf("expected menu squelch to be true, got %v", mi.MenuItems.Utilities.Squelch)
	}

	if mi.MenuItems.Utilities.PasswordAndLock != true {
		t.Errorf("expected menu password and lock to be true, got %v", mi.MenuItems.Utilities.PasswordAndLock)
	}

	if mi.MenuItems.Utilities.Vox != false {
		t.Errorf("expected menu vox to be false, got %v", mi.MenuItems.Utilities.Vox)
	}

	if mi.MenuItems.Utilities.DisplayMode != true {
		t.Errorf("expected menu display mode to be true, got %v", mi.MenuItems.Utilities.DisplayMode)
	}

	if mi.MenuItems.Utilities.ProgramRadio != true {
		t.Errorf("expected menu program radio to be true, got %v", mi.MenuItems.Utilities.ProgramRadio)
	}

}
