package encoding

import (
	"go/types"
)

type Decoder interface {
	GetEntityID() string
	GetEntityType() types.BasicKind
	Decode(buf []byte, base uint32) interface{}
}
