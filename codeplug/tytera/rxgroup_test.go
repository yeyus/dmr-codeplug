package tytera

import (
	"encoding/json"
	"fmt"
	//"github.com/yeyus/dmr-codeplug/proto/tytera"
	"io/ioutil"
	"log"
	"testing"
)

type rxGroupsTest struct{}

func (rxGroupsTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestRxGroupsParsing(t *testing.T) {
	d := rxGroupsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	rg := GetRxGroupListGroup()

	b, err := json.MarshalIndent(rg.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", rg.Groups)
}

func TestRxGroupsProto(t *testing.T) {
	d := rxGroupsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	rg := GetRxGroupListGroup()
	rg.Decode(content[:], 0x125)

	if len(rg.Groups.Entries) != 250 {
		t.Errorf("expected number of contacts to be 250, but got %d", len(rg.Groups.Entries))
	}

	testRxGroupEntry(t, rg, 0, "DMR", []uint32{952, 953, 951})
}

func testRxGroupEntry(t *testing.T, r RxGroupListGroup, idx int, name string, contacts []uint32) {
	entry := r.Groups.Entries[idx]
	if entry.Name != name {
		t.Errorf("[idx %d] expected rxgroup name to be %s, but got %s", idx, name, entry.Name)
	}

	diff := difference(contacts, entry.ContactIndex)
	if len(diff) > 0 {
		t.Errorf("[idx %d] expected rxgroup groups to be the same, difference is %v", idx, diff)
	}
}

func difference(a, b []uint32) (diff []uint32) {
	m := make(map[uint32]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
