package tytera

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// DFU Image constants
var SZ_TARGET_SIGNATURE = [...]byte{'T', 'a', 'r', 'g', 'e', 't'}

type DfuSeTarget struct {
	AlternateSetting uint8
	Named            bool
	Name             string
	Size             uint32
	Element          []DfuSeElement
}

func ParseTargetsInImage(file []byte, offset uint32, numTargets uint8) (result []DfuSeTarget, err error) {

	var i uint8
	for i = 0; i < numTargets; i++ {
		target := DfuSeTarget{}

		targetSignature := file[offset+0 : offset+6]
		if !bytes.Equal(targetSignature, SZ_TARGET_SIGNATURE[:]) {
			return result, fmt.Errorf("DfuSe Target: invalid target signature 0x%x at offset %x, expected 0x%x", targetSignature, offset, SZ_TARGET_SIGNATURE)
		}

		target.AlternateSetting = file[offset+6]
		target.Named = binary.LittleEndian.Uint32(file[offset+7:offset+11]) != 0
		targetName := file[offset+11 : offset+266]
		target.Name = string(targetName[:strings.Index(string(targetName), "\x00")])
		target.Size = binary.LittleEndian.Uint32(file[offset+266 : offset+270])

		// Elements parsing
		numImages := binary.LittleEndian.Uint32(file[offset+270 : offset+274])
		elements, err := ParseElementsInTarget(file, offset+274, numImages)
		target.Element = elements
		if err != nil {
			return result, err
		}

		offset = offset + 274 + target.Size
		result = append(result, target)
	}

	return result, err
}

func (t *DfuSeTarget) String() string {
	var elements bytes.Buffer
	for _, element := range t.Element {
		elements.WriteString(element.String())
	}

	return fmt.Sprintf(
		"[> Alternate: 0x%X Name: %s Size: 0x%X\n"+
			"\t--- Elements ---\n"+
			"%s"+
			"\t--- /Elements ---\n",
		t.AlternateSetting, t.Name, t.Size, elements.String())
}
