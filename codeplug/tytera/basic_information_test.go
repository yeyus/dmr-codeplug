package tytera

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

type biTest struct{}

func (biTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestBasicInformationParsing(t *testing.T) {
	d := biTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	bi := GetBasicInformationDecoder()

	b, err := json.MarshalIndent(bi.Decode(content[:], 0x125), "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

	fmt.Printf("%+v", bi.BasicInformation)
}
