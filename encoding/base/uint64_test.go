package base

import (
	"go/types"
	"testing"
)

func TestUint64DecoderLE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0x12, 0}

	d := &Uint64Decoder{
		EntityID:  "com.test.uint64",
		Offset:    3,
		Length:    5,
		Endianess: LittleEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.uint64" {
		t.Errorf("expected id to be com.test.bcd, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint64 {
		t.Errorf("returned type for decoder is not uint64, got %s", d.GetEntityType())
	}

	if o != uint64(0x1234567812) {
		t.Errorf("expected decoder output to be 0x1234567812, but got 0x%x", o)
	}

}

func TestUint64DecoderBE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0x12, 0}

	d := &Uint64Decoder{
		EntityID:  "com.test.uint64",
		Offset:    3,
		Length:    5,
		Endianess: BigEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.uint64" {
		t.Errorf("expected id to be com.test.bcd, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Uint64 {
		t.Errorf("returned type for decoder is not uint64, got %s", d.GetEntityType())
	}

	if o != uint64(0x1278563412) {
		t.Errorf("expected decoder output to be 0x12378563412, but got 0x%x", o)
	}

}

func TestUint64DecoderExceedLength(t *testing.T) {
	data := [...]byte{0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0x12, 0}

	d := &Uint64Decoder{
		EntityID:  "com.test.uint64",
		Offset:    3,
		Length:    9,
		Endianess: LittleEndian,
	}

	o := d.Decode(data[:], 0)

	if o != nil {
		t.Errorf("expected decoder to fail if length > 8 bytes, but got %d", o)
	}
}
