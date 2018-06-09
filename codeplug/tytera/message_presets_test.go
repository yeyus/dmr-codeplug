package tytera

import (
	"encoding/json"
	"fmt"
	//"github.com/yeyus/dmr-codeplug/proto/tytera"
	"io/ioutil"
	"log"
	"testing"
)

type messagesTest struct{}

func (messagesTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestMessagePresetsParsing(t *testing.T) {
	d := messagesTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	mp := GetMessagePresets()

	m, err := json.MarshalIndent(mp.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(m))
}

func TestMessagePresetsProto(t *testing.T) {
	d := messagesTest{}
	content := d.getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	mp := GetMessagePresets()

	mp.Decode(content[:], 0x125)

	messages := mp.Messages.Messages

	if len(messages) != 50 {
		t.Errorf("number of messages presets should be 50, found %d", len(messages))
	}

	if messages[0] != "Hello" {
		t.Errorf("first message should be \"Hello\", but got \"%s\"", messages[0])
	}

	if messages[1] != "This is some sample message" {
		t.Errorf("first message should be \"This is some sample message\", but got \"%s\"", messages[1])
	}

}
