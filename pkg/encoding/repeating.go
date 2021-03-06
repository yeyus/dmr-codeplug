package encoding

import (
	"go/types"
)

type IsSetPredicate func(interface{}) bool

type RepeatedDecoder struct {
	Offset       uint32
	RecordLength uint32
	Records      uint32
	Decoder      Decoder
	Predicate    IsSetPredicate
}

func (t *RepeatedDecoder) GetEntityID() string {
	return t.Decoder.GetEntityID()
}

func (t *RepeatedDecoder) GetEntityType() types.BasicKind {
	return types.UnsafePointer
}

func (t *RepeatedDecoder) Decode(buf []byte, base uint32) interface{} {
	if base+t.Offset > uint32(len(buf)) || base+t.Offset+(t.Records*t.RecordLength) > uint32(len(buf)) {
		return nil
	}

	var arr []interface{}
	for i := 0; i < int(t.Records); i++ {
		start := base + t.Offset + (uint32(i) * t.RecordLength)
		obj := t.Decoder.Decode(buf, start)
		if t.Predicate == nil || t.Predicate(obj) {
			arr = append(arr, obj)
		}
	}

	return arr
}
