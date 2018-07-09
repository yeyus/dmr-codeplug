package base

import (
	"go/types"
	"testing"
)

func TestBitDecoder1(t *testing.T) {
	data := [...]byte{0, 0x10, 0}

	d := &BitDecoder{
		EntityID:  "com.test.bit",
		Offset:    1,
		BitOffset: 4,
	}

	o := d.Decode(data[:], 0).(bool)

	if d.GetEntityID() != "com.test.bit" {
		t.Errorf("expected id to be com.test.bit, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Bool {
		t.Errorf("returned type for decoder is not bool, got %v", d.GetEntityType())
	}

	if !o {
		t.Errorf("expected decoder output to be true, but got %t", o)
	}
}

func TestBitDecoder2(t *testing.T) {
	data := [...]byte{0, 0x10, 0}

	d := &BitDecoder{
		EntityID:  "com.test.bit",
		Offset:    1,
		BitOffset: 5,
	}

	o := d.Decode(data[:], 0).(bool)

	if d.GetEntityID() != "com.test.bit" {
		t.Errorf("expected id to be com.test.bit, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Bool {
		t.Errorf("returned type for decoder is not bool, got %v", d.GetEntityType())
	}

	if o {
		t.Errorf("expected decoder output to be false, but got %t", o)
	}
}
