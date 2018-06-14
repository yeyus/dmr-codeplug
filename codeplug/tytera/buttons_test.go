package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"io/ioutil"
	"log"
	"testing"
)

type buttonsTest struct{}

func (buttonsTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestButtonsParsing(t *testing.T) {
	d := buttonsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	bs := GetButtonsGroup()

	b, err := json.MarshalIndent(bs.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))

	fmt.Printf("%+v", bs.Buttons)
}

func TestButtonsProto(t *testing.T) {
	d := buttonsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	bs := GetButtonsGroup()

	bs.Decode(content[:], 0x125)

	if bs.Buttons.LongPressDuration != 1000 {
		t.Errorf("expected buttons long press duration to be 1000, got %d", bs.Buttons.LongPressDuration)
	}

	if bs.Buttons.SideShort_1 != tytera.ButtonFunction_HIGH_LOW_POWER {
		t.Errorf("expected side button 1 short press function to be unassigned, got %d", bs.Buttons.SideShort_1)
	}

	if bs.Buttons.SideShort_2 != tytera.ButtonFunction_SCAN_ONOFF {
		t.Errorf("expected side button 2 short press function to be unassigned, got %d", bs.Buttons.SideShort_2)
	}

	if bs.Buttons.SideLong_1 != tytera.ButtonFunction_ALL_ALERT_TONES_ONOFF {
		t.Errorf("expected side button 1 long press function to be unassigned, got %d", bs.Buttons.SideLong_1)
	}

	if bs.Buttons.SideLong_2 != tytera.ButtonFunction_MANUAL_DIAL {
		t.Errorf("expected side button 2 long press function to be unassigned, got %d", bs.Buttons.SideLong_2)
	}

	// one touch access
	if len(bs.Buttons.OneTouchAccess) != 6 {
		t.Errorf("expected one touch access size to be 6, got %d", len(bs.Buttons.OneTouchAccess))
	}

	// [0] should be DIGITAL CALL to contact 11
	if bs.Buttons.OneTouchAccess[0].Mode != tytera.SystemType_DIGITAL {
		t.Errorf("expected one touch access 0 mode to be DIGITAL, got %v", bs.Buttons.OneTouchAccess[0].Mode)
	}

	if bs.Buttons.OneTouchAccess[0].CallType != tytera.OneTouchCallType_CALL {
		t.Errorf("expected one touch access 0 call type to be CALL, got %v", bs.Buttons.OneTouchAccess[0].CallType)
	}

	if bs.Buttons.OneTouchAccess[0].ContactIndex != 11 {
		t.Errorf("expected one touch access 0 contact index to be 11, got %v", bs.Buttons.OneTouchAccess[0].ContactIndex)
	}

	if bs.Buttons.OneTouchAccess[0].MessagePreset != 0 {
		t.Errorf("expected one touch access 0 message preset to be zero value, got %v", bs.Buttons.OneTouchAccess[0].MessagePreset)
	}

	if bs.Buttons.OneTouchAccess[0].DtmfEncodePreset != 0 {
		t.Errorf("expected one touch access 0 dtmf encode preset to be zero value, got %v", bs.Buttons.OneTouchAccess[0].DtmfEncodePreset)
	}

	// [3] should be DIGITAL MESSAGE preset 1 to contact 6
	if bs.Buttons.OneTouchAccess[3].Mode != tytera.SystemType_DIGITAL {
		t.Errorf("expected one touch access 0 mode to be DIGITAL, got %v", bs.Buttons.OneTouchAccess[3].Mode)
	}

	if bs.Buttons.OneTouchAccess[3].CallType != tytera.OneTouchCallType_TEXT_MESSAGE {
		t.Errorf("expected one touch access 3 call type to be TEXT_MESSAGE, got %v", bs.Buttons.OneTouchAccess[3].CallType)
	}

	if bs.Buttons.OneTouchAccess[3].ContactIndex != 6 {
		t.Errorf("expected one touch access 3 contact index to be 6, got %v", bs.Buttons.OneTouchAccess[3].ContactIndex)
	}

	if bs.Buttons.OneTouchAccess[3].MessagePreset != 1 {
		t.Errorf("expected one touch access 3 message preset to be 1, got %v", bs.Buttons.OneTouchAccess[3].MessagePreset)
	}

	if bs.Buttons.OneTouchAccess[3].DtmfEncodePreset != 0 {
		t.Errorf("expected one touch access 3 dtmf encode preset to be zero value, got %v", bs.Buttons.OneTouchAccess[3].DtmfEncodePreset)
	}

	// [5] should be ANALOG DTMF2 using preset 1
	if bs.Buttons.OneTouchAccess[5].Mode != tytera.SystemType_ANALOG {
		t.Errorf("expected one touch access 5 mode to be ANALOG, got %v", bs.Buttons.OneTouchAccess[5].Mode)
	}

	if bs.Buttons.OneTouchAccess[5].CallType != tytera.OneTouchCallType_DTMF2 {
		t.Errorf("expected one touch access 5 call type to be DTMF2, got %v", bs.Buttons.OneTouchAccess[5].CallType)
	}

	if bs.Buttons.OneTouchAccess[5].ContactIndex != 0 {
		t.Errorf("expected one touch access 5 contact index to be zero value, got %v", bs.Buttons.OneTouchAccess[5].ContactIndex)
	}

	if bs.Buttons.OneTouchAccess[5].MessagePreset != 0 {
		t.Errorf("expected one touch access 5 message preset to be zero value, got %v", bs.Buttons.OneTouchAccess[5].MessagePreset)
	}

	if bs.Buttons.OneTouchAccess[5].DtmfEncodePreset != 1 {
		t.Errorf("expected one touch access 5 dtmf encode preset to be 1, got %v", bs.Buttons.OneTouchAccess[5].DtmfEncodePreset)
	}

	// contact keys
	if bs.Buttons.ContactKey_0 != 10 {
		t.Errorf("expected contact key 0 to be 10, got %d", bs.Buttons.ContactKey_0)
	}

	if bs.Buttons.ContactKey_1 != 4 {
		t.Errorf("expected contact key 1 to be 4, got %d", bs.Buttons.ContactKey_1)
	}

	if bs.Buttons.ContactKey_2 != 0 {
		t.Errorf("expected contact key 2 to be zero value, got %d", bs.Buttons.ContactKey_2)
	}

	if bs.Buttons.ContactKey_9 != 13 {
		t.Errorf("expected contact key 9 to be 13, got %d", bs.Buttons.ContactKey_9)
	}

}
