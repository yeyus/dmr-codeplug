package base

import (
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
	"testing"
)

func TestBCDToneDecoder_None(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0xFF, 0xFF, 0, 0}

	d := &BCDToneDecoder{
		EntityID: "com.test.tone",
		Offset:   4,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.tone" {
		t.Errorf("expected id to be com.test.tone but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.UnsafePointer {
		t.Errorf("expected type for decoder is not unsafePointer, got %v", d.GetEntityType())
	}

	tone := o.(tytera.Tone)
	if tone.Type != tytera.ToneSystem_NO_TONE {
		t.Errorf("expected tone system to be not set, but got %s", tone.Type)
	}

	if tone.Frequency != 0 {
		t.Errorf("expected tone frequency to be 0, but got %d", tone.Frequency)
	}

	if tone.Code != 0 {
		t.Errorf("expected tone code to be 0, but got %d", tone.Code)
	}

}

func TestBCDToneDecoder_CTCSS(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0x48, 0x11, 0, 0}

	d := &BCDToneDecoder{
		EntityID: "com.test.tone",
		Offset:   4,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.tone" {
		t.Errorf("expected id to be com.test.tone but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.UnsafePointer {
		t.Errorf("expected type for decoder is not unsafePointer, got %v", d.GetEntityType())
	}

	tone := o.(tytera.Tone)
	if tone.Type != tytera.ToneSystem_CTCSS {
		t.Errorf("expected tone system to be CTCSS, but got %s", tone.Type)
	}

	if tone.Frequency != 1148 {
		t.Errorf("expected tone frequency to be 1148, but got %d", tone.Frequency)
	}

	if tone.Code != 0 {
		t.Errorf("expected tone code to be 0, but got %d", tone.Code)
	}

}

func TestBCDToneDecoder_DCSN(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0x47, 0x80, 0, 0}

	d := &BCDToneDecoder{
		EntityID: "com.test.tone",
		Offset:   4,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.tone" {
		t.Errorf("expected id to be com.test.tone but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.UnsafePointer {
		t.Errorf("expected type for decoder is not unsafePointer, got %v", d.GetEntityType())
	}

	tone := o.(tytera.Tone)
	if tone.Type != tytera.ToneSystem_DCS_NORMAL {
		t.Errorf("expected tone system to be DCS NORMAL, but got %s", tone.Type)
	}

	if tone.Frequency != 0 {
		t.Errorf("expected tone frequency to be 0, but got %d", tone.Frequency)
	}

	if tone.Code != 47 {
		t.Errorf("expected tone code to be 47, but got %d", tone.Code)
	}

}

func TestBCDToneDecoder_DCSI(t *testing.T) {
	data := [...]byte{0, 0, 0, 0, 0x25, 0xC3, 0, 0}

	d := &BCDToneDecoder{
		EntityID: "com.test.tone",
		Offset:   4,
	}

	o := d.Decode(data[:], 0)

	if d.GetEntityID() != "com.test.tone" {
		t.Errorf("expected id to be com.test.tone but got %s", d.GetEntityID())
	}

	if d.GetEntityType() != types.UnsafePointer {
		t.Errorf("expected type for decoder is not unsafePointer, got %v", d.GetEntityType())
	}

	tone := o.(tytera.Tone)
	if tone.Type != tytera.ToneSystem_DCS_INVERTED {
		t.Errorf("expected tone system to be DCS INVERTED, but got %s", tone.Type)
	}

	if tone.Frequency != 0 {
		t.Errorf("expected tone frequency to be 0, but got %d", tone.Frequency)
	}

	if tone.Code != 325 {
		t.Errorf("expected tone code to be 325, but got %d", tone.Code)
	}

}
