package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestScanListParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	sl := GetScanListGroup()

	b, err := json.MarshalIndent(sl.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", sl.ScanLists)
}

func TestScanListsProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	sl := GetScanListGroup()
	sl.Decode(content[:], 0x125)

	if len(sl.ScanLists.Entries) != 2 {
		t.Errorf("expected number of scan lists to be 2, but got %d", len(sl.ScanLists.Entries))
	}

	testScanListEntry(t, sl, 0, "Scanlist", 0, 0, 0xFFFF, 500, 2000, []uint32{})
	testScanListEntry(t, sl, 1, "Unit Test", 22, 2, 76, 700, 2750, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31})
}

func testScanListEntry(t *testing.T, r ScanListGroup, idx int, name string, pc1 uint32, pc2 uint32, tx uint32, holdTime uint32, sampleTime uint32, channels []uint32) {
	entry := r.ScanLists.Entries[idx]
	if entry.Name != name {
		t.Errorf("[idx %d] expected scan list name to be %s, but got %s", idx, name, entry.Name)
	}

	diff := difference(channels, entry.ChannelIndex)
	if len(diff) > 0 {
		t.Errorf("[idx %d] expected scan list channels to be the same, difference is %v", idx, diff)
	}
}
