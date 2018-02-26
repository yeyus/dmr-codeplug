package base

import (
	"go/types"
	"testing"
)

func TestUTF16StringDecoderLE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0x41, 0, 0x42, 0, 0x43, 0, 0x44, 0, 0, 0}

	d := &UTF16StringDecoder{
		EntityID:  "com.test.utf16",
		Offset:    4,
		Length:    10,
		Endianess: LittleEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.utf16" {
		t.Errorf("expected id to be com.test.utf16, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.String {
		t.Errorf("returned type for decoder is not string, got %s", d.GetEntityType())
	}

	if o != "ABCD" {
		t.Errorf("expected decoder output to be ABCD, but got %s", o)
	}

}

func TestUTF16StringDecoderBE(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0x41, 0, 0x42, 0, 0x43, 0, 0x44, 0, 0, 0}

	d := &UTF16StringDecoder{
		EntityID:  "com.test.utf16",
		Offset:    3,
		Length:    10,
		Endianess: BigEndian,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.utf16" {
		t.Errorf("expected id to be com.test.utf16, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.String {
		t.Errorf("returned type for decoder is not string, got %s", d.GetEntityType())
	}

	if o != "ABCD" {
		t.Errorf("expected decoder output to be ABCD, but got %s", o)
	}
}
