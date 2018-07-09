package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"reflect"
	"testing"
)

func TestDTMFParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	dg := GetDTMFGroup()

	g, err := json.MarshalIndent(dg.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(g))
	fmt.Printf("%+v", dg.DTMF)
}

func TestDTMFSystemProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	dg := GetDTMFGroup()
	dg.Decode(content[:], 0x125)

	if len(dg.DTMF.Systems) != 4 {
		t.Errorf("expected systems length to be 4, but got %d", len(dg.DTMF.Systems))
	}

	compareSystemEntry(t, dg, 0, tytera.DTMFSystemEntry{
		FirstDigitDelay:   400,
		FirstDigitTime:    150,
		DigitDurationTime: 140,
		DigitIntervalTime: 130,
		StarHashDigitTime: 120,
		DKeyAssignment:    110,
		NextSequence:      2,
		AutoResetTime:     10,
		SideTone:          true,
		PttId:             tytera.PttIdType_PRE_ONLY,
		GroupCode:         tytera.GroupCodeType_D,
		KeyUpEncode:       "ABCD1",
		KeyDownEncode:     "1234",
	})

	compareSystemEntry(t, dg, 1, tytera.DTMFSystemEntry{
		FirstDigitDelay:   400,
		FirstDigitTime:    100,
		DigitDurationTime: 100,
		DigitIntervalTime: 100,
		StarHashDigitTime: 100,
		DKeyAssignment:    100,
		NextSequence:      2,
		AutoResetTime:     10,
		SideTone:          true,
	})

	compareSystemEntry(t, dg, 2, tytera.DTMFSystemEntry{
		FirstDigitDelay:   370,
		FirstDigitTime:    80,
		DigitDurationTime: 80,
		DigitIntervalTime: 60,
		StarHashDigitTime: 0,
		DKeyAssignment:    0,
		NextSequence:      4,
		AutoResetTime:     12,
		SideTone:          false,
		PttId:             tytera.PttIdType_PRE_AND_POST,
		GroupCode:         tytera.GroupCodeType_HASH,
		KeyUpEncode:       "123456",
		KeyDownEncode:     "4567",
	})

}

func compareSystemEntry(t *testing.T, dg *DTMFGroup, idx int, x tytera.DTMFSystemEntry) {
	var a tytera.DTMFSystemEntry = *dg.DTMF.Systems[idx]

	if !reflect.DeepEqual(a, x) {
		t.Errorf("expected dtmf system index %d to be equal to => \n Expected: (type %T)\n %+v\n Got: (type %T)\n %+v", idx, x, x, a, a)
	}

}

func TestDTMFEncodesProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	dg := GetDTMFGroup()
	dg.Decode(content[:], 0x125)

	if len(dg.DTMF.Encodes) != 32 {
		t.Errorf("expected DTMF encodes to be 32, but got %d", len(dg.DTMF.Encodes))
	}

	if dg.DTMF.Encodes[0] != "ABCD*#" {
		t.Errorf("expected DTMF encode 0 to be ABCD*#, but got %s", dg.DTMF.Encodes[0])
	}

	if dg.DTMF.Encodes[1] != "12345" {
		t.Errorf("expected DTMF encode 1 to be 12345, but got %s", dg.DTMF.Encodes[1])
	}

	if dg.DTMF.Encodes[4] != "D*ADB**#" {
		t.Errorf("expected DTMF encode 4 to be D*ADB**#, but got %s", dg.DTMF.Encodes[4])
	}

}

func TestDTMFDecodeProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/04_Oct_2017.rdt")

	dg := GetDTMFGroup()
	dg.Decode(content[:], 0x125)

	if len(dg.DTMF.Decodes) != 8 {
		t.Errorf("expected DTMF decodes to be 8, but got %d", len(dg.DTMF.Decodes))
	}

	compareDecodeEntry(t, dg, 0, tytera.DTMFDecodeEntry{
		DtmfId:             "",
		ResponseType:       tytera.ResponseType_GENERAL,
		DecodeType:         tytera.DecodeType_DECODE_TYPE_NONE,
		AckEncodeIndex:     0,
		AckDelayTime:       0,
		RevertChannelIndex: 0,
	})

	compareDecodeEntry(t, dg, 1, tytera.DTMFDecodeEntry{
		DtmfId:             "12345*#",
		ResponseType:       tytera.ResponseType_KILL,
		DecodeType:         tytera.DecodeType_SEL_CALL,
		AckEncodeIndex:     2,
		AckDelayTime:       150,
		RevertChannelIndex: 9,
	})
}

func compareDecodeEntry(t *testing.T, dg *DTMFGroup, idx int, x tytera.DTMFDecodeEntry) {
	var a tytera.DTMFDecodeEntry = *dg.DTMF.Decodes[idx]

	if !reflect.DeepEqual(a, x) {
		t.Errorf("expected dtmf decode index %d to be equal to => \n Expected: (type %T)\n %+v\n Got: (type %T)\n %+v", idx, x, x, a, a)
	}

}
