package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrivacySettingsParsing(t *testing.T) {
	content := getRDTBytes("../../../test/usa_codeplug.rdt")

	ps := GetPrivacySettingsGroup()

	p, err := json.MarshalIndent(ps.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(p))
}

func TestPrivacySettingsProto(t *testing.T) {
	content := getRDTBytes("../../../test/usa_codeplug.rdt")

	ps := GetPrivacySettingsGroup()

	ps.Decode(content[:], 0x125)

	if len(ps.Privacy.BasicKeys) != 16 {
		t.Errorf("Basic Keys length expected to be 16, got %d", len(ps.Privacy.BasicKeys))
	}

	if len(ps.Privacy.EnhancedKeys) != 8 {
		t.Errorf("Enhanced  Keys length expected to be 8, got %d", len(ps.Privacy.EnhancedKeys))
	}

	if fmt.Sprintf("%X", ps.Privacy.BasicKeys[0]) != fmt.Sprintf("%X", [...]byte{0xAA, 0xAA}) {
		t.Errorf("expected basic key 0 to be 0xAAAA, got 0x%s", fmt.Sprintf("%X", ps.Privacy.BasicKeys[0]))
	}

	if fmt.Sprintf("%X", ps.Privacy.BasicKeys[1]) != fmt.Sprintf("%X", [...]byte{0xFF, 0xFF}) {
		t.Errorf("expected basic key 1 to be 0xFFFF, got 0x%s", fmt.Sprintf("%X", ps.Privacy.BasicKeys[1]))
	}

	if fmt.Sprintf("%X", ps.Privacy.BasicKeys[15]) != fmt.Sprintf("%X", [...]byte{0xBB, 0xBB}) {
		t.Errorf("expected basic key 1 to be 0xBBBB, got 0x%s", fmt.Sprintf("%X", ps.Privacy.BasicKeys[15]))
	}

	if fmt.Sprintf("%X", ps.Privacy.EnhancedKeys[0]) != fmt.Sprintf("%X", [...]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xDE, 0xAD, 0xBE, 0xEF}) {
		t.Errorf("expected enhanced key 0 to be 0xFFFFFFFFFFFFFFFFFFFFFFFFDEADBEEF, got 0x%s", fmt.Sprintf("%X", ps.Privacy.EnhancedKeys[0]))
	}

	if fmt.Sprintf("%X", ps.Privacy.EnhancedKeys[7]) != fmt.Sprintf("%X", [...]byte{0xFE, 0xED, 0xBE, 0xEF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}) {
		t.Errorf("expected enhanced key 7 to be 0xFFFFFFFFFFFFFFFFFFFFFFFFDEADBEEF, got 0x%s", fmt.Sprintf("%X", ps.Privacy.EnhancedKeys[7]))
	}

}
