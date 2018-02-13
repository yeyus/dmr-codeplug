package tytera

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
)

// DFU Prefix constants
var DFU_HEADER = [...]byte{'D', 'f', 'u', 'S', 'e'}

const B_VERSION uint8 = 1

// DFU Suffix constants
var UC_DFU_SIGNATURE = [...]byte{'U', 'F', 'D'}
var DFU_SPEC_NUMBER = [...]byte{0x1A, 0x01}

const DFU_SUFFIX_LENGHT uint8 = 16
const TYTERA_BROKEN_CRC uint32 = 0x8E657112

type TyteraDfuSe struct {
	Data      []byte
	Version   uint8
	Size      uint32
	Target    []DfuSeTarget
	VersionID uint16
	ProductID uint16
	VendorID  uint16
	CRC       uint32
}

func GetTyteraDfuSe(file []byte) (TyteraDfuSe, error) {
	t := TyteraDfuSe{}
	t.Data = file

	// parse header
	err := t.parseHeader(file)
	if err != nil {
		return t, err
	}

	// parse suffix
	err = t.parseSuffix(file)
	if err != nil {
		return t, err
	}

	numTargets := file[10]
	targets, err := ParseTargetsInImage(file, 0xB, numTargets)
	t.Target = targets
	if err != nil {
		return t, err
	}

	return t, nil
}

func (t *TyteraDfuSe) parseHeader(file []byte) error {
	// Check file header
	if !bytes.Equal(file[0:5], DFU_HEADER[:]) {
		return errors.New("file header did not match")
	}

	// Check version
	if file[5] != B_VERSION {
		return errors.New("DfuSe version mismatch")
	}
	t.Version = file[5]

	// Parse size
	t.Size = binary.LittleEndian.Uint32(file[6:11])

	return nil
}

func (t *TyteraDfuSe) parseSuffix(file []byte) error {
	// Find signature and determine if tytera's cps is adding and additional
	// 0x100 bytes to the image before the suffix
	var suffixOffset uint32 = 0
	if bytes.Equal(file[t.Size+8:t.Size+0x10+1], UC_DFU_SIGNATURE[:]) {
		suffixOffset = t.Size
	} else if bytes.Equal(file[t.Size+0x100+8:t.Size+0x100+10+1], UC_DFU_SIGNATURE[:]) {
		suffixOffset = t.Size + 0x100
	} else {
		return errors.New("DfuSe DFU Suffix signature not found")
	}

	t.VersionID = binary.LittleEndian.Uint16(file[suffixOffset+0 : suffixOffset+2])
	t.ProductID = binary.LittleEndian.Uint16(file[suffixOffset+2 : suffixOffset+4])
	t.VendorID = binary.LittleEndian.Uint16(file[suffixOffset+4 : suffixOffset+6])

	if !bytes.Equal(file[suffixOffset+6:suffixOffset+8], DFU_SPEC_NUMBER[:]) {
		return errors.New("DfuSe invalid DFU specification signature")
	}

	if file[suffixOffset+11] != 0x10 {
		return errors.New("DfuSe invalid DFU suffix length")
	}

	fileCRC := binary.LittleEndian.Uint32(file[suffixOffset+12 : suffixOffset+16])
	if fileCRC == TYTERA_BROKEN_CRC {
		// Tytera uses a fixed CRC value
		t.CRC = TYTERA_BROKEN_CRC
	} else {
		// If other CRC is present it will be check against calculated CRC
		calculatedCRC := crc32.ChecksumIEEE(file[:len(file)-4])
		if calculatedCRC != fileCRC {
			return fmt.Errorf("DfuSe file checksum 0x%x doesn't match dfuse checksum 0x%x", fileCRC, calculatedCRC)
		}
	}

	return nil
}

func (t *TyteraDfuSe) String() string {
	var targets bytes.Buffer
	for _, target := range t.Target {
		targets.WriteString(target.String())
	}

	return fmt.Sprintf(
		"\n==== Image ==== \n"+
			"Tytera DfuSe Image."+
			"Version: 0x%X Size: 0x%X CRC: 0x%X\n"+
			"VersionID: 0x%X ProductID: 0x%X VendorID: 0x%X \n"+
			"==== Targets ====\n"+
			"%s\n"+
			"==== /Image ==== \n", t.Version, t.Size, t.CRC, t.VersionID, t.ProductID, t.VendorID, targets.String())
}
