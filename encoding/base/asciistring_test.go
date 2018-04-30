package base

import (
	"go/types"
	"testing"
)

func TestASCIIStringDecoder(t *testing.T) {
	data := [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 'T', 'E', 'S', 'T', 0, '0', '0'}

	d := &ASCIIStringDecoder{
		EntityID: "com.test.string",
		Offset:   8,
		Length:   7,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.string" {
		t.Errorf("expected id to be com.test.string, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.String {
		t.Errorf("returned type for decoder is not string, got %v", d.GetEntityType())
	}

	if o != "TEST" {
		t.Errorf("expected decoder output to be TEST, but got %s", o)
	}

	data = [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 0, 'T', 'E', 'S', 'T', '0', '0'}
	o = d.Decode(data[:], 0)

	if o != "" {
		t.Errorf("expected decoder output to be \"\", but got %s", o)
	}

	data = [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 'T', 'E', 'E', 'E', 'E', 'S', 'T'}
	o = d.Decode(data[:], 0)

	if o != "TEEEEST" {
		t.Errorf("expected decoder output to be TEEEEST, but got %s", o)
	}
}

func TestASCIIStringDecoderOutOfBounds(t *testing.T) {
	data := [...]byte{'T', 'E', 'S', 'T'}

	d := &ASCIIStringDecoder{
		EntityID: "com.test.string",
		Offset:   8,
		Length:   7,
	}

	o := d.Decode(data[:], 0)

	if o != nil {
		t.Errorf("expected to return nil if out of bounds")
	}
}

func TestASCIIStringDecoderNull(t *testing.T) {
	d := &ASCIIStringDecoder{
		EntityID: "com.test.string",
		Offset:   8,
		Length:   7,
	}

	o := d.Decode(nil, 0)

	if o != nil {
		t.Errorf("expected to return nil if out of bounds")
	}

}
