package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBasicInformationParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	bi := GetBasicInformationGroup()

	b, err := json.MarshalIndent(bi.Decode(content[:], 0x125), "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

	fmt.Printf("%+v", bi.BasicInformation)
}
