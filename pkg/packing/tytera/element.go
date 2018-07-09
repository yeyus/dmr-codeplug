package tytera

import (
	"encoding/binary"
	"fmt"
)

type DfuSeElement struct {
	Start uint32
	Size  uint32
	Data  []byte
}

func ParseElementsInTarget(file []byte, offset uint32, numElements uint32) (result []DfuSeElement, err error) {

	var i uint32
	for i = 0; i < numElements; i++ {
		image := DfuSeElement{}

		image.Start = binary.LittleEndian.Uint32(file[offset+0 : offset+4])
		image.Size = binary.LittleEndian.Uint32(file[offset+4 : offset+8])
		if uint32(len(file)) < offset+8+image.Size {
			return result, fmt.Errorf("DfuSe Image: image size mismatch")
		}
		image.Data = file[offset+8 : offset+8+image.Size]

		offset = offset + 8 + image.Size

		result = append(result, image)
	}

	return result, err
}

func (e *DfuSeElement) String() string {
	return fmt.Sprintf("\tELEMENT -> Start: 0x%X Size: 0x%X\n", e.Start, e.Size)
}
