package tytera

import (
	//	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMD380CodeplugProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	plug := GetMD380Codeplug()
	p := plug.Decode(content[:], 0x125)

	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Printf("Error: %s \n", err.Error())
	}

	fmt.Printf("%s", b)
}
