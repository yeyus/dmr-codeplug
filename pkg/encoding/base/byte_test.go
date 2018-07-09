package base

import (
	"go/types"
	"testing"
)

func TestByteDecoder(t *testing.T) {
	data := [...]byte{0, 0xAA, 0xBB, 0xCC, 0xDD}

	d := &ByteDecoder{
		EntityID: "com.test.byte",
		Offset:   0,
	}

	o := d.Decode(data[:], 1)

	if d.GetEntityID() != "com.test.byte" {
		t.Errorf("expected id to be com.test.byte, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint8 {
		t.Errorf("returned type for decoder is not uint8, got %v", d.GetEntityType())
	}

	if o != uint8(0xAA) {
		t.Errorf("expected decoder output to be 0xAA, but got %d", o)
	}

	o = d.Decode(data[:], 2)

	if o != uint8(0xBB) {
		t.Errorf("expected decoder output to be 0xBB, but got %d", o)
	}
}
