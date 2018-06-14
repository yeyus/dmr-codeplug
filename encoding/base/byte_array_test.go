package base

import (
	"fmt"
	"go/types"
	"testing"
)

func TestByteArrayDecoder(t *testing.T) {
	data := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	d := &ByteArrayDecoder{
		EntityID: "com.test.byteArray",
		Offset:   3,
		Length:   5,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.byteArray" {
		t.Errorf("expected id to be com.test.byteArray, but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.Byte {
		t.Errorf("expected type for decoder is not Byte, got %v", d.GetEntityType())
	}

	fmt.Printf("%+v %T", o, o)
	for i, b := range o.([]byte)[:] {
		if b != data[i+3] {
			t.Errorf("expected array value to be 0x%X but got 0x%X", b, data[i+3])
			break
		}
	}
}
