package packing

import (
	"bytes"
	"errors"
)

// DFU Prefix constants
const DFU_HEADER = [...]byte{'D', 'f', 'u', 'S', 'e'}
const B_VERSION uint8 = 1

// DFU Suffix constants
const UC_DFU_SIGNATURE = [...]byte{'U', 'F', 'D'}
const DFU_SPEC_NUMBER = [...]byte{0x1A, 0x01}
const DFU_SUFFIX_LENGHT uint8 = 16
const TYTERA_BROKEN_CRC uint32 = 0x127165BE

// DFU Image constants
const SZ_TARGET_SIGNATURE = [...]byte{'T', 'a', 'r', 'g', 'e', 't'}

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
	Named            boolean
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

	/*
	   Parse DFU Header
	*/

	// Check file header
	if !bytes.Equal(file[0:5], DFU_HEADER) {
		return _, errors.New("file header did not match")
	}

	// Check DfuSe version
	if file[5] != B_VERSION {
		return _, errors.New("DfuSe version mismatch")
	}

	var size uint32 = binary.BigEndian.Uint32(file[6:11])

	/*
	   Parse DFU Suffix
	*/

	// Find signature and determine if tytera's cps is adding and additional 0x100 bytes
	// to the image before the suffix
	var tyterasExtraBytes boolean = false
	if bytes.Equal(file[size+8:size+10+1], UC_DFU_SIGNATURE) {
		tyterasExtraBytes = false
	} else if bytes.Equal(file[size+0x100+8:size+0x100+10+1], UC_DFU_SIGNATURE) {
		tyterasExtraBytes = true
	} else {
		return _, errors.New("DfuSe DFU Suffix signature not found")
	}

}

func (t *TyteraDfuSe) Validate() {

}

func (t *TyteraDfuSe) GetData() {

}
