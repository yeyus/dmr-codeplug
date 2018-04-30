package base

import (
	"go/types"
	"testing"
)

func TestBCDDecoderLE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0, 0}

	d := &BCDDecoder{
		EntityID:  "com.test.bcd",
		Offset:    3,
		Length:    4,
		Endianess: LittleEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.bcd" {
		t.Errorf("expected id to be com.test.bcd, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint64 {
		t.Errorf("returned type for decoder is not uint64, got %v", d.GetEntityType())
	}

	if o != uint64(12345678) {
		t.Errorf("expected decoder output to be 12345678, but got %d", o)
	}

}

func TestBCDDecoderBE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0x78, 0x56, 0x34, 0x12, 0, 0}

	d := &BCDDecoder{
		EntityID:  "com.test.bcd",
		Offset:    3,
		Length:    4,
		Endianess: BigEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.bcd" {
		t.Errorf("expected id to be com.test.bcd, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint64 {
		t.Errorf("returned type for decoder is not uint64, got %v", d.GetEntityType())
	}

	if o != uint64(12345678) {
		t.Errorf("expected decoder output to be 12345678, but got %d", o)
	}

}
