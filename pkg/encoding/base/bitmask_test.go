package base

import (
	"go/types"
	"testing"
)

func TestBitMaskDecoderMeta(t *testing.T) {
	data := [...]byte{0, 0xF0, 0x0F, 0x3C, 0x00}

	d := &BitMaskDecoder{
		EntityID: "com.test.bitmask",
		Offset:   0,
		BitMask:  0xF0,
	}

	o := d.Decode(data[:], 1)

	if d.GetEntityID() != "com.test.bitmask" {
		t.Errorf("expected id to be com.test.bitmask but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint8 {
		t.Errorf("returned type for decoder is not uint8, got %v", d.GetEntityType())
	}

	if o != uint8(0x0F) {
		t.Errorf("expected decoder output to be 0x0F, but got %d", o)
	}

	o = d.Decode(data[:], 2)

	if o != uint8(0x00) {
		t.Errorf("expected decoder output to be 0x00, but got %d", o)
	}

	o = d.Decode(data[:], 3)

	if o != uint8(0x03) {
		t.Errorf("expected decoder output to be 0x03, but got %d", o)
	}
}

func TestBitMaskDecoderMove(t *testing.T) {
	data := [...]byte{0xC0}

	d := &BitMaskDecoder{
		EntityID: "com.test.bitmask",
		Offset:   0,
		BitMask:  0xC0,
	}

	o := d.Decode(data[:], 0)

	if o != uint8(0x03) {
		t.Errorf("expected decoder output to be 0x03, but got %d", o)
	}
}
