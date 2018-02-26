package tytera

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

type gsTest struct{}

func (gsTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestGeneralSettingsParsing(t *testing.T) {
	d := gsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	gs := GetGeneralSettingsGroup()

	g, err := json.MarshalIndent(gs.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(g))

	fmt.Printf("%+v", gs.GeneralSettings)
}
