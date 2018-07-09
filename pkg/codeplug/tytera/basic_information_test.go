package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"testing"
)

func TestBasicInformationParsing(t *testing.T) {
	content := getRDTBytes("../../../test/04_Oct_2017.rdt")

	bi := GetBasicInformationGroup()

	b, err := json.MarshalIndent(bi.Decode(content[:], 0x125), "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

	fmt.Printf("%+v", bi.BasicInformation)
}

func TestBasicInformationProto(t *testing.T) {
	content := getRDTBytes("../../../test/04_Oct_2017.rdt")

	bi := GetBasicInformationGroup()
	bi.Decode(content[:], 0x125)

	if bi.BasicInformation.ModelName != "MD390" {
		t.Errorf("expected model name to be MD390 but got %s", bi.BasicInformation.ModelName)
	}

	if len(bi.BasicInformation.RadioBands) != 2 {
		t.Errorf("expected radio bands length to be 1, but got %d", len(bi.BasicInformation.RadioBands))
	}

	if bi.BasicInformation.RadioBands[0] != tytera.RadioBand_UHF_400_480 {
		t.Errorf("expected radio band 0 to be 400-480Mhz but got, %s", bi.BasicInformation.RadioBands[0])
	}

	if bi.BasicInformation.RadioBands[1] != tytera.RadioBand_VHF_136_174 {
		t.Errorf("expected radio band 1 to be DISABLED, but got %s", bi.BasicInformation.RadioBands[1])
	}

	if bi.BasicInformation.HardwareVersion != "V01.00" {
		t.Errorf("expected hardware version to be V01.00, but got %s", bi.BasicInformation.HardwareVersion)
	}

	if bi.BasicInformation.McuVersion != "S013.020" {
		t.Errorf("expected mcu version to be S013.020, but got %s", bi.BasicInformation.McuVersion)
	}

	if bi.BasicInformation.DeviceId != "583845Q0021479948357" {
		t.Errorf("expected device id to be 583845Q0021479948357, but got %s", bi.BasicInformation.DeviceId)
	}

	if bi.BasicInformation.CpsVersion != "V01.36" {
		t.Errorf("expected cps version to be V01.36, but got %s", bi.BasicInformation.CpsVersion)
	}
}
