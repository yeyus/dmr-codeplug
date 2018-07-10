package tytera

import (
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"reflect"
	"testing"
)

func TestGPSProto(t *testing.T) {
	content := getRDTBytes("../../../test/04_Oct_2017.rdt")

	gs := GetGPSGroup()
	gs.Decode(content[:], 0x125)

	if len(gs.GPS.Entries) != 16 {
		t.Errorf("expected number of gps systems to be 16, but got %d", len(gs.GPS.Entries))
	}

	compareGPSSystem(t, gs, 0, tytera.GPSSystemEntry{
		RevertChannelIndex:      16,
		ReportInterval:          120,
		DestinationContactIndex: 11,
	})

	compareGPSSystem(t, gs, 1, tytera.GPSSystemEntry{
		RevertChannelIndex:      0,
		ReportInterval:          300,
		DestinationContactIndex: 11,
	})
}

func compareGPSSystem(t *testing.T, g *GPSGroup, idx int, x tytera.GPSSystemEntry) {
	var a tytera.GPSSystemEntry = *g.GPS.Entries[idx]

	if !reflect.DeepEqual(a, x) {
		t.Errorf("expected gps system index %d to be equal to => \n Expected: (type %T)\n %+v\n Got: (type %T)\n %+v", idx, x, x, a, a)
	}
}
