package pgpwords

import (
	"encoding/hex"
	"testing"
)

func TestEncodeWords(t *testing.T) {

	var tst16 = []struct {
		in  uint16
		out string
	}{
		{0xE582, "topmost Istanbul"},
		{0x82E5, "miser travesty"},
		{0xFFFF, "Zulu Yucatán"},
	}

	for _, tt := range tst16 {
		if enc := Encode16(tt.in); enc != tt.out {
			t.Errorf("Encode16(%s) failed: got '%v' expected '%v'", tt.in, enc, tt.out)
		}
	}

	var tst32 = []struct {
		in  uint32
		out string
	}{
		{0xE58282E5, "topmost Istanbul miser travesty"},
		{0xFFFFFFFF, "Zulu Yucatán Zulu Yucatán"},
	}

	for _, tt := range tst32 {
		if enc := Encode32(tt.in); enc != tt.out {
			t.Errorf("Encode32(%s) failed: got '%v' expected '%v'", tt.in, enc, tt.out)
		}
	}

}

func TestEncode(t *testing.T) {
	in, _ := hex.DecodeString("E58294F2E9A22748")
	out := `topmost Istanbul Pluto vagabond treadmill Pacific brackish dictator`

	if enc := Encode(in); enc != out {
		t.Errorf("Encode(%s) failed: got '%v' expected '%v'", enc, out)
	}
}
