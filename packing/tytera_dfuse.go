package packing

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"strings"
)

// DFU Prefix constants
var DFU_HEADER = [...]byte{'D', 'f', 'u', 'S', 'e'}

const B_VERSION uint8 = 1

// DFU Suffix constants
var UC_DFU_SIGNATURE = [...]byte{'U', 'F', 'D'}
var DFU_SPEC_NUMBER = [...]byte{0x1A, 0x01}

const DFU_SUFFIX_LENGHT uint8 = 16
const TYTERA_BROKEN_CRC uint32 = 0x8E657112

// DFU Image constants
var SZ_TARGET_SIGNATURE = [...]byte{'T', 'a', 'r', 'g', 'e', 't'}

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

type DfuSeTarget struct {
	AlternateSetting uint8
	Named            bool
	Name             string
	Size             uint32
	Element          []DfuSeElement
}

type DfuSeElement struct {
	Start uint32
	Size  uint32
	Data  []byte
}

func GetTyteraDfuSe(file []byte) (TyteraDfuSe, error) {

	dfuse := TyteraDfuSe{}

	/*
	   Parse DFU Header
	*/

	// Check file header
	if !bytes.Equal(file[0:5], DFU_HEADER[:]) {
		return dfuse, errors.New("file header did not match")
	}

	// Check version
	if file[5] != B_VERSION {
		return dfuse, errors.New("DfuSe version mismatch")
	}
	dfuse.Version = file[5]

	// Parse size
	var size uint32 = binary.LittleEndian.Uint32(file[6:11])
	dfuse.Size = size

	/*
	   Parse DFU Suffix
	*/

	// Find signature and determine if tytera's cps is adding and additional 0x100 bytes
	// to the image before the suffix
	var suffixOffset uint32 = 0
	if bytes.Equal(file[size+8:size+0x10+1], UC_DFU_SIGNATURE[:]) {
		suffixOffset = size
	} else if bytes.Equal(file[size+0x100+8:size+0x100+10+1], UC_DFU_SIGNATURE[:]) {
		suffixOffset = size + 0x100
	} else {
		return dfuse, errors.New("DfuSe DFU Suffix signature not found")
	}

	dfuse.VersionID = binary.LittleEndian.Uint16(file[suffixOffset+0 : suffixOffset+2])
	dfuse.ProductID = binary.LittleEndian.Uint16(file[suffixOffset+2 : suffixOffset+4])
	dfuse.VendorID = binary.LittleEndian.Uint16(file[suffixOffset+4 : suffixOffset+6])

	if !bytes.Equal(file[suffixOffset+6:suffixOffset+8], DFU_SPEC_NUMBER[:]) {
		return dfuse, errors.New("DfuSe invalid DFU specification signature")
	}

	if file[suffixOffset+11] != 0x10 {
		return dfuse, errors.New("DfuSe invalid DFU suffix length")
	}

	fileCRC := binary.LittleEndian.Uint32(file[suffixOffset+12 : suffixOffset+16])
	if fileCRC == TYTERA_BROKEN_CRC {
		// Tytera uses a fixed CRC value
		dfuse.CRC = TYTERA_BROKEN_CRC
	} else {
		// If other CRC is present it will be check against calculated CRC
		calculatedCRC := crc32.ChecksumIEEE(file[:len(file)-4])
		if calculatedCRC != fileCRC {
			return dfuse, fmt.Errorf("DfuSe file checksum 0x%x doesn't match dfuse checksum 0x%x", fileCRC, calculatedCRC)
		}
	}

	err := dfuse.parseTargets(file)
	if err != nil {
		return dfuse, err
	}

	return dfuse, nil
}

func (t *TyteraDfuSe) parseTargets(file []byte) error {

	var numTargets uint8 = file[10]
	if numTargets == 0 {
		return nil
	}

	return t.parseTargetRecursive(file, 0xB, numTargets)
}

func (t *TyteraDfuSe) parseTargetRecursive(file []byte, offset uint32, remaining uint8) error {
	target := DfuSeTarget{}

	targetSignature := file[offset+0 : offset+6]
	if !bytes.Equal(targetSignature, SZ_TARGET_SIGNATURE[:]) {
		return fmt.Errorf("DfuSe Target: invalid target signature 0x%x at offset %x, expected 0x%x", targetSignature, offset, SZ_TARGET_SIGNATURE)
	}

	target.AlternateSetting = file[offset+6]
	target.Named = binary.LittleEndian.Uint32(file[offset+7:offset+11]) != 0
	targetName := file[offset+11 : offset+266]
	target.Name = string(targetName[:strings.Index(string(targetName), "\x00")])
	target.Size = binary.LittleEndian.Uint32(file[offset+266 : offset+270])

	//numImages := binary.LittleEndian.Uint32(file[offset+270 : offset+274])
	// TODO process images in targets

	t.Target = append(t.Target, target)

	remaining--
	if remaining > 0 {
		return t.parseTargetRecursive(file, offset+274+target.Size, remaining)
	}

	return nil
}

func (t *TyteraDfuSe) Validate() {

}

func (t *TyteraDfuSe) GetData() {

}
