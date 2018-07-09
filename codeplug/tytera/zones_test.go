package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestZonesParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	rg := GetZonesGroup()

	b, err := json.MarshalIndent(rg.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", rg.Zones)
}

func TestZonesProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	rg := GetZonesGroup()
	rg.Decode(content[:], 0x125)

	if len(rg.Zones.Entries) != 12 {
		t.Errorf("expected number of zones to be 12, but got %d", len(rg.Zones.Entries))
	}

	testZoneEntry(t, rg, 0, "Global", []uint32{1, 2, 3, 4, 5, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74})
	testZoneEntry(t, rg, 1, "BM Regional", []uint32{5, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 67, 68, 69, 70, 71})
	testZoneEntry(t, rg, 2, "BM Region 0", []uint32{2, 26, 14, 15, 7, 53, 59, 48, 40, 10, 76, 77, 78, 67, 68, 69})
	testZoneEntry(t, rg, 3, "BM Region 1", []uint32{2, 27, 42, 61, 49, 44, 11, 57, 76, 77, 78, 67, 68, 69, 70, 71})
	testZoneEntry(t, rg, 4, "BM Region 2", []uint32{2, 28, 50, 51, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74, 75})
	testZoneEntry(t, rg, 5, "BM Region 3", []uint32{2, 29, 12, 43, 56, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74})
	testZoneEntry(t, rg, 6, "BM Region 4", []uint32{2, 30, 16, 17, 18, 52, 58, 13, 62, 41, 76, 77, 78, 67, 68, 69})
	testZoneEntry(t, rg, 7, "BM Region 5", []uint32{2, 31, 6, 46, 8, 19, 25, 9, 76, 77, 78, 67, 68, 69, 70, 71})
	testZoneEntry(t, rg, 8, "BM Region 6", []uint32{2, 32, 22, 20, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74, 75})
	testZoneEntry(t, rg, 9, "BM Region 7", []uint32{2, 33, 24, 23, 60, 55, 63, 66, 47, 37, 21, 76, 77, 78, 67, 68})
	testZoneEntry(t, rg, 10, "BM Region 8", []uint32{2, 34, 45, 54, 64, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74})
	testZoneEntry(t, rg, 11, "BM Region 9", []uint32{2, 35, 65, 38, 39, 76, 77, 78, 67, 68, 69, 70, 71, 72, 73, 74})
}

func testZoneEntry(t *testing.T, r *ZonesGroup, idx int, name string, channels []uint32) {
	entry := r.Zones.Entries[idx]
	if entry.Name != name {
		t.Errorf("[idx %d] expected zone name to be %s, but got %s", idx, name, entry.Name)
	}

	diff := difference(channels, entry.ChannelIndex)
	if len(diff) > 0 {
		t.Errorf("[idx %d] expected zone channels to be the same, difference is %v", idx, diff)
	}
}
