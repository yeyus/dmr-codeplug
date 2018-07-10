package tytera

import (
	"io/ioutil"
	"log"
	"testing"
)

type dfuseTest struct{}

func (dfuseTest) getRDTBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestMD380_GetTyteraDfuSe(t *testing.T) {
	d := dfuseTest{}
	content := d.getRDTBytes("../../../test/usa_codeplug.rdt")

	dfuse, err := GetTyteraDfuSe(content)
	if err != nil {
		t.Fatal(err)
	}

	if dfuse.Version != 1 {
		t.Errorf("DFU version expected 1 but got %d", dfuse.Version)
	}

	if dfuse.Size != 0x40125 {
		t.Errorf("DFU size expected 0x40125 but got %x", dfuse.Size)
	}

	if dfuse.VersionID != 0x0200 {
		t.Errorf("DFU version id expected 0x0200 but got %x", dfuse.VersionID)
	}

	if dfuse.ProductID != 0xDF11 {
		t.Errorf("DFU product id expected 0xDF11 but got %x", dfuse.ProductID)
	}

	if dfuse.VendorID != 0x0483 {
		t.Errorf("DFU vendor id expected 0x0483 but got %x", dfuse.VendorID)
	}

	if dfuse.CRC != TYTERA_BROKEN_CRC {
		t.Errorf("DFU checksum is invalid, expected 0x%x but got 0x%x", TYTERA_BROKEN_CRC, dfuse.CRC)
	}

	t.Logf("%s\n", dfuse.String())
}

func TestMD390_GetTyteraDfuSe(t *testing.T) {
	d := dfuseTest{}
	content := d.getRDTBytes("../../../test/04_Oct_2017.rdt")

	dfuse, err := GetTyteraDfuSe(content)
	if err != nil {
		t.Fatal(err)
	}

	if dfuse.Version != 1 {
		t.Errorf("DFU version expected 1 but got %d", dfuse.Version)
	}

	if dfuse.Size != 0x40125 {
		t.Errorf("DFU size expected 0x40125 but got %x", dfuse.Size)
	}

	if dfuse.VersionID != 0x0200 {
		t.Errorf("DFU version id expected 0x0200 but got %x", dfuse.VersionID)
	}

	if dfuse.ProductID != 0xDF11 {
		t.Errorf("DFU product id expected 0xDF11 but got %x", dfuse.ProductID)
	}

	if dfuse.VendorID != 0x0483 {
		t.Errorf("DFU vendor id expected 0x0483 but got %x", dfuse.VendorID)
	}

	if dfuse.CRC != TYTERA_BROKEN_CRC {
		t.Errorf("DFU checksum is invalid, expected 0x%x but got 0x%x", TYTERA_BROKEN_CRC, dfuse.CRC)
	}

	t.Logf("%s\n", dfuse.String())
}

func TestMD2017_GetTyteraDfuSe(t *testing.T) {
	d := dfuseTest{}
	content := d.getRDTBytes("../../../test/MD2017-G0NEF-UK.rdt")

	dfuse, err := GetTyteraDfuSe(content)
	if err != nil {
		t.Fatal(err)
	}

	if dfuse.Version != 1 {
		t.Errorf("DFU version expected 1 but got %d", dfuse.Version)
	}

	if dfuse.Size != 0x40125 {
		t.Errorf("DFU size expected 0x40125 but got %x", dfuse.Size)
	}

	if dfuse.VersionID != 0x0200 {
		t.Errorf("DFU version id expected 0x0200 but got %x", dfuse.VersionID)
	}

	if dfuse.ProductID != 0xDF11 {
		t.Errorf("DFU product id expected 0xDF11 but got %x", dfuse.ProductID)
	}

	if dfuse.VendorID != 0x0483 {
		t.Errorf("DFU vendor id expected 0x0483 but got %x", dfuse.VendorID)
	}

	if dfuse.CRC != TYTERA_BROKEN_CRC {
		t.Errorf("DFU checksum is invalid, expected 0x%x but got 0x%x", TYTERA_BROKEN_CRC, dfuse.CRC)
	}

	t.Logf("%s\n", dfuse.String())
}

func TestGetCodeplugs(t *testing.T) {
	d := dfuseTest{}
	content := d.getRDTBytes("../../../test/usa_codeplug.rdt")

	dfuse, err := GetTyteraDfuSe(content)
	if err != nil {
		t.Fatal(err)
	}

	codeplugs := dfuse.GetCodeplugs()

	if len(codeplugs) != 1 {
		t.Errorf("expected 1 codeplug but got %d", len(codeplugs))
	}

	if len(codeplugs[0]) != 0x40000 {
		t.Errorf("expected codeplug to have size 0x40000 but got %d", len(codeplugs[0]))
	}
}
