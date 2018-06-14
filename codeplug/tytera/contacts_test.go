package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"io/ioutil"
	"log"
	"testing"
)

type contactsTest struct{}

func (contactsTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestContactsParsing(t *testing.T) {
	d := contactsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	cs := GetContactsGroup()

	b, err := json.MarshalIndent(cs.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", cs.Contacts)
}

func TestContactsProto(t *testing.T) {
	d := contactsTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	cs := GetContactsGroup()
	cs.Decode(content[:], 0x125)

	testContact(t, cs, 0, 2, "Local", true, tytera.CallType_GROUP_CALL)
	testContact(t, cs, 1, 9, "Hotspot", false, tytera.CallType_PRIVATE_CALL)
	testContact(t, cs, 2, 0xFFFFFF, "Worldwide", false, tytera.CallType_ALL_CALL)
	testContact(t, cs, 3, 93, "North America", false, tytera.CallType_GROUP_CALL)
	testContact(t, cs, 40, 3139, "Ohio", false, tytera.CallType_GROUP_CALL)
	testContact(t, cs, 77, 312, "TAC 312", false, tytera.CallType_GROUP_CALL)

}

func testContact(t *testing.T, c ContactsGroup, idx int, id uint32, name string, callTone bool, callType tytera.CallType) {
	entry := c.Contacts.Entries[idx]
	if entry.Id != id {
		t.Errorf("[idx %d] expected contact id to be %d, but got %d", idx, id, entry.Id)
	}

	if entry.Name != name {
		t.Errorf("[idx %d] expected contact name to be %s, but got %s", idx, name, entry.Name)
	}

	if entry.CallReceiveTone != callTone {
		t.Errorf("[idx %d] expected call tone to be %t, but got %t", idx, entry.CallReceiveTone, callTone)
	}

	if entry.CallType != callType {
		t.Errorf("[idx %d] expected call type to be %s, but got %s", idx, entry.CallType, callType)
	}
}
