package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"testing"
)

func TestGeneralSettingsParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	gs := GetGeneralSettingsGroup()

	g, err := json.MarshalIndent(gs.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(g))

	fmt.Printf("%+v", gs.GeneralSettings)
}

// TODO Some comprehensive testing

func TestGeneralSettingsProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	gs := GetGeneralSettingsGroup()
	gs.Decode(content[:], 0x125)

	if gs.GeneralSettings.ScreenLines[0] != "MD-380G" {
		t.Errorf("expected screen line 0 to be MD-380G, but got %s", gs.GeneralSettings.ScreenLines[0])
	}

	if gs.GeneralSettings.ScreenLines[1] != "AG6VW" {
		t.Errorf("expected screen line 1 to be AG6VW, but got %s", gs.GeneralSettings.ScreenLines[1])
	}

	if gs.GeneralSettings.MonitorType != tytera.MonitorType_OPEN_SQUELCH {
		t.Errorf("expected monitor type to be OPEN_SQUELCH, but got %s", gs.GeneralSettings.MonitorType)
	}

	if gs.GeneralSettings.DisableLeds {
		t.Errorf("expected disable leds to be false, but got %t", gs.GeneralSettings.DisableLeds)
	}

	if gs.GeneralSettings.TalkPermitTone != tytera.SystemType_NONE {
		t.Errorf("expected talk permit tone to be NONE, but got %s", gs.GeneralSettings.TalkPermitTone)
	}

	if !gs.GeneralSettings.DisablePasswordAndLock {
		t.Errorf("expected disable password and lock to be true, but got %t", gs.GeneralSettings.DisablePasswordAndLock)
	}

	if !gs.GeneralSettings.DisableChFreeTone {
		t.Errorf("expected disable ch free tone to be true, but got %t", gs.GeneralSettings.DisableChFreeTone)
	}

	if !gs.GeneralSettings.DisableTones {
		t.Errorf("expected disable tones to be true, but got %t", gs.GeneralSettings.DisableTones)
	}

	if !gs.GeneralSettings.BatSaveReceiveMode {
		t.Errorf("expected bat save receive mode to be true, but got %t", gs.GeneralSettings.BatSaveReceiveMode)
	}

	if !gs.GeneralSettings.BatSavePreamble {
		t.Errorf("expected bat save preamble to be true, but got %t", gs.GeneralSettings.BatSavePreamble)
	}

	if gs.GeneralSettings.IntroScreenMode != tytera.IntroScreenMode_CHAR_STRING {
		t.Errorf("expected intro screen mode to be CHAR_STRING but got, %s", gs.GeneralSettings.IntroScreenMode)
	}

	if gs.GeneralSettings.RadioId != 1106167 {
		t.Errorf("expected radio id to be 1106167 but got %d", gs.GeneralSettings.RadioId)
	}

	if gs.GeneralSettings.TxPreamble != 300 {
		t.Errorf("expected tx preamble to be 300 ms but got %d", gs.GeneralSettings.TxPreamble)
	}

	if gs.GeneralSettings.GroupCallHang != 3000 {
		t.Errorf("expected group call hang to be 3000, but got %d", gs.GeneralSettings.GroupCallHang)
	}

	if gs.GeneralSettings.PrivateCallHang != 4000 {
		t.Errorf("expected private call hang to be 4000, but got %d", gs.GeneralSettings.PrivateCallHang)
	}

	if gs.GeneralSettings.VoxSensitivity != 3 {
		t.Errorf("expected vox sensitivity to be 3, but got %d", gs.GeneralSettings.VoxSensitivity)
	}

	if gs.GeneralSettings.RxLowBattery != 120 {
		t.Errorf("expected rx low battery to be 120 but got %d", gs.GeneralSettings.RxLowBattery)
	}

	if gs.GeneralSettings.CallAlertTone != 0 {
		t.Errorf("expected call alert tone to be 0 but got %d", gs.GeneralSettings.CallAlertTone)
	}

	if gs.GeneralSettings.LoneWorkerResponse != 1 {
		t.Errorf("expected lone worker response to be 1 but got %d", gs.GeneralSettings.LoneWorkerResponse)
	}

	if gs.GeneralSettings.LoneWorkerReminder != 10 {
		t.Errorf("expected lone worker reminder to be 10, but got %d", gs.GeneralSettings.LoneWorkerReminder)
	}

	if gs.GeneralSettings.ScanDigitalHangTime != 1000 {
		t.Errorf("expected scan digital hang time to be 1000, but got %d", gs.GeneralSettings.ScanDigitalHangTime)
	}

	if gs.GeneralSettings.ScanAnalogHangTime != 1000 {
		t.Errorf("expected scan analog hang time to be 1000, but got %d", gs.GeneralSettings.ScanAnalogHangTime)
	}

	if gs.GeneralSettings.BacklightTimeout != 5 {
		t.Errorf("expected backlight timeout to be 5, but got %d", gs.GeneralSettings.BacklightTimeout)
	}

	if gs.GeneralSettings.KeypadLock != 255 {
		t.Errorf("expected keypad lock to be 255 (manual), but got %d", gs.GeneralSettings.KeypadLock)
	}

	if gs.GeneralSettings.OperationMode != tytera.OperationMode_CHANNEL {
		t.Errorf("expected operation mode to be CHANNEL, but got %s", gs.GeneralSettings.OperationMode)
	}

	if gs.GeneralSettings.PowerOnPassword != 0 {
		t.Errorf("expected power on password to be 0, but got %d", gs.GeneralSettings.PowerOnPassword)
	}

	if gs.GeneralSettings.RadioProgrammingPassword != 0 {
		t.Errorf("expected radio programming password to be 0, but got %d", gs.GeneralSettings.RadioProgrammingPassword)
	}

	if gs.GeneralSettings.PcProgrammingPassword != "" {
		t.Errorf("expected pc programming password to be \"\", but got %s", gs.GeneralSettings.PcProgrammingPassword)
	}

	if gs.GeneralSettings.RadioName != "AG6VW" {
		t.Errorf("expected radio name to be AG6VW, but got %s", gs.GeneralSettings.RadioName)
	}
}
