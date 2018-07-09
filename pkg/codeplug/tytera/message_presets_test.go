package tytera

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessagePresetsParsing(t *testing.T) {
	content := getRDTBytes("../../../test/usa_codeplug.rdt")

	mp := GetMessagePresetsGroup()

	m, err := json.MarshalIndent(mp.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(m))
}

func TestMessagePresetsProto(t *testing.T) {
	content := getRDTBytes("../../../test/usa_codeplug.rdt")

	mp := GetMessagePresetsGroup()

	mp.Decode(content[:], 0x125)

	messages := mp.Messages.Messages

	if len(messages) != 2 {
		t.Errorf("number of messages presets should be 2, found %d", len(messages))
	}

	if messages[0] != "Hello" {
		t.Errorf("first message should be \"Hello\", but got \"%s\"", messages[0])
	}

	if messages[1] != "This is some sample message" {
		t.Errorf("first message should be \"This is some sample message\", but got \"%s\"", messages[1])
	}

}
