package tytera

import (
	"encoding/json"
	"fmt"
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
