package base

import (
	"go/types"
	"testing"
)

func TestDTMFString(t *testing.T) {
	data := [...]byte{0, 0, 0x1E, 0x21, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED, 0x0F, 0x21, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED}

	d := &DTMFStringDecoder{
		EntityID: "com.test.dtmfString",
		Offset:   2,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.dtmfString" {
		t.Errorf("expected id to be com.test.dtmfString, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.String {
		t.Errorf("expected type to be String, but got %v", d.GetEntityType())
	}

	if o != "123456789ABCD*#0123456789ABCD*" {
		t.Errorf("expected string to be 123456789ABCD*#0123456789ABCD*, but got %s", o)
	}
}

func TestDTMFString_2(t *testing.T) {
	data := [...]byte{0, 0, 0x0F, 0x21, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED, 0x0F, 0x21, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED}

	d := &DTMFStringDecoder{
		EntityID: "com.test.dtmfString",
		Offset:   2,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.dtmfString" {
		t.Errorf("expected id to be com.test.dtmfString, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.String {
		t.Errorf("expected type to be String, but got %v", d.GetEntityType())
	}

	if o != "123456789ABCD*#" {
		t.Errorf("expected string to be 123456789ABCD*#, but got %s", o)
	}
}
