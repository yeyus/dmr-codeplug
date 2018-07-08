package tytera

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/yeyus/dmr-codeplug/encoding"
	"github.com/yeyus/dmr-codeplug/encoding/base"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"go/types"
	"strconv"
	"time"
)

type BasicInformationGroup struct {
	EntityID         string
	Base             uint32
	Length           uint32
	Decoders         []encoding.Decoder
	BasicInformation tytera.BasicInformation
}

func GetBasicInformationGroup() *BasicInformationGroup {
	b := BasicInformationGroup{
		EntityID:         "com.tytera.basic",
		Base:             0x210B,
		Length:           0, // ?
		BasicInformation: tytera.BasicInformation{},
	}

	b.Decoders = []encoding.Decoder{
		&base.ASCIIStringDecoder{
			EntityID: "com.tytera.basic.modelName",
			Offset:   0,
			Length:   16,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.basic.band[0].index",
			Offset:   0x11,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.basic.band[1].index",
			Offset:   0x12,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.band[0].lowerLimit",
			Offset:    0x14,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.band[0].upperLimit",
			Offset:    0x16,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.band[1].lowerLimit",
			Offset:    0x18,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.band[1].upperLimit",
			Offset:    0x1B,
			Length:    2,
			Endianess: base.BigEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.hardwareVersion",
			Offset:    0x34,
			Length:    4,
			Endianess: base.BigEndian,
		},
		&base.ByteDecoder{
			EntityID: "com.tytera.basic.mcu.variant",
			Offset:   0x40,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.mcu.version",
			Offset:    0x41,
			Length:    3,
			Endianess: base.LittleEndian,
		},
		&base.Uint64Decoder{
			EntityID:  "com.tytera.basic.deviceID.udid1",
			Offset:    0x48,
			Length:    5,
			Endianess: base.BigEndian,
		},
		&base.ASCIIStringDecoder{
			EntityID: "com.tytera.basic.deviceID.udid2",
			Offset:   0x4D,
			Length:   7,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.lastProgramDate",
			Offset:    0x2101,
			Length:    7,
			Endianess: base.LittleEndian,
		},
		&base.BCDDecoder{
			EntityID:  "com.tytera.basic.cpsVersion",
			Offset:    0x2108,
			Length:    4,
			Endianess: base.LittleEndian,
		},
	}

	return &b
}

func (t *BasicInformationGroup) GetEntityID() string {
	return t.EntityID
}

func (t *BasicInformationGroup) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *BasicInformationGroup) Decode(buf []byte, base uint32) interface{} {
	for _, d := range t.Decoders {
		t.mapValue(d, buf, base)
	}

	return t.BasicInformation
}

func (t *BasicInformationGroup) mapValue(d encoding.Decoder, buf []byte, base uint32) {
	switch id := d.GetEntityID(); id {
	case "com.tytera.basic.modelName":
		t.BasicInformation.ModelName = d.Decode(buf, base).(string)
	case "com.tytera.basic.band[0].index",
		"com.tytera.basic.band[1].index":
		o := d.Decode(buf, base).(uint8)
		t.BasicInformation.RadioBands = append(t.BasicInformation.RadioBands, t.mapRadioBand(o))
	case "com.tytera.basic.hardwareVersion":
		var o uint64 = d.Decode(buf, base).(uint64)
		// if value was FF FF FF FF
		if o == 166666665 {
			t.BasicInformation.HardwareVersion = ""
		} else {
			t.BasicInformation.HardwareVersion = fmt.Sprintf("V%02d.%02d", o/100, o%100)
		}
	case "com.tytera.basic.mcu.variant":
		var o uint8 = d.Decode(buf, base).(uint8)
		t.BasicInformation.McuVersion = fmt.Sprintf("%c", o)
	case "com.tytera.basic.mcu.version":
		var o uint64 = d.Decode(buf, base).(uint64)
		// if value was FF FF FF FF
		if o == 166666665 {
			t.BasicInformation.McuVersion = ""
		} else {
			t.BasicInformation.McuVersion = fmt.Sprintf("%s%03d.%03d", t.BasicInformation.McuVersion, o/1000, o%1000)
		}
	case "com.tytera.basic.deviceID.udid1":
		var o uint64 = d.Decode(buf, base).(uint64)
		// if value was FF FF FF FF
		if o == 166666665 {
			t.BasicInformation.DeviceId = ""
		} else {
			t.BasicInformation.DeviceId = fmt.Sprintf("%013d", o)
		}
	case "com.tytera.basic.deviceID.udid2":
		var o string = d.Decode(buf, base).(string)
		// if value was FF FF FF FF
		t.BasicInformation.DeviceId = fmt.Sprintf("%s%s", reverse(o), t.BasicInformation.DeviceId)
	case "com.tytera.basic.lastProgramDate":
		var o uint64 = d.Decode(buf, base).(uint64)
		date := strconv.FormatUint(o, 10)
		seconds := date[12:14]
		minutes := date[10:12]
		hours := date[8:10]
		day := date[6:8]
		month := date[4:6]
		year := date[0:4]
		dateStr := fmt.Sprintf("%s-%s-%sT%s:%s:%s", year, month, day, hours, minutes, seconds)
		fmt.Print(dateStr)
		tm, err := time.Parse("2006-01-02T15:04:05", dateStr)
		if err != nil {
			t.BasicInformation.LastProgramDate = nil
			fmt.Print(err)
		} else {
			s := int64(tm.Unix())
			t.BasicInformation.LastProgramDate = &timestamp.Timestamp{Seconds: s, Nanos: 0}

		}
	case "com.tytera.basic.cpsVersion":
		var o uint64 = d.Decode(buf, base).(uint64)
		t.BasicInformation.CpsVersion = fmt.Sprintf("V%1d%1d.%1d%1d", (o/1000000)%10, (o/10000)%10, (o/100)%10, o%10)
	}

}

func (t *BasicInformationGroup) mapRadioBand(radioBandIndex uint8) (b tytera.RadioBand) {
	switch x := radioBandIndex; x {
	case 0:
		b = tytera.RadioBand_VHF_136_174
	case 1:
		b = tytera.RadioBand_UHF_350_400
	case 2:
		b = tytera.RadioBand_UHF_400_480
	case 3:
		b = tytera.RadioBand_UHF_450_520
	case 255:
		b = tytera.RadioBand_DISABLED
	}

	return
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
