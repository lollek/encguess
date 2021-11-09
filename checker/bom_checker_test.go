package checker

import (
	"github.com/lollek/encodingutil/encoding"
	"testing"
)

func TestUTF1(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xF7)
	checker.CheckNext(0x64)
	checker.CheckNext(0x4C)
	if !checker.Validates() || checker.Encoding() != encoding.UTF1 {
		t.Fatalf("Failed to handle UTF-1 file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF1 {
		t.Fatalf("Failed to handle UTF-1 file")
	}
}

func TestUTF7(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0x2B)
	checker.CheckNext(0x2F)
	checker.CheckNext(0x76)
	if !checker.Validates() || checker.Encoding() != encoding.UTF7 {
		t.Fatalf("Failed to handle UTF-7 file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF7 {
		t.Fatalf("Failed to handle UTF-7 file")
	}
}

func TestUTF8(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xEF)
	checker.CheckNext(0xBB)
	checker.CheckNext(0xBF)
	if !checker.Validates() || checker.Encoding() != encoding.UTF8 {
		t.Fatalf("Failed to handle UTF-8 file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF8 {
		t.Fatalf("Failed to handle UTF-8 file")
	}
}

func TestUTF16_BE(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xFE)
	checker.CheckNext(0xFF)
	if !checker.Validates() || checker.Encoding() != encoding.UTF16_BE {
		t.Fatalf("Failed to handle UTF-16 BE file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF16_BE {
		t.Fatalf("Failed to handle UTF-16 BE file")
	}
}

func TestUTF16_LE(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xFF)
	checker.CheckNext(0xFE)
	if !checker.Validates() || checker.Encoding() != encoding.UTF16_LE {
		t.Fatalf("Failed to handle UTF-16 LE file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF16_LE {
		t.Fatalf("Failed to handle UTF-16 LE file")
	}
}

func TestUTF32_BE(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0x00)
	checker.CheckNext(0x00)
	checker.CheckNext(0xFE)
	checker.CheckNext(0xFF)
	if !checker.Validates() || checker.Encoding() != encoding.UTF32_BE {
		t.Fatalf("Failed to handle UTF-32 BE file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF32_BE {
		t.Fatalf("Failed to handle UTF-32 BE file")
	}
}

func TestUTF32_LE(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xFF)
	checker.CheckNext(0xFE)
	checker.CheckNext(0x00)
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF32_LE {
		t.Fatalf("Failed to handle UTF-32 LE file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF32_LE {
		t.Fatalf("Failed to handle UTF-32 LE file")
	}
}

func TestUTF_EBCDIC(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xDD)
	checker.CheckNext(0x73)
	checker.CheckNext(0x66)
	checker.CheckNext(0x73)
	if !checker.Validates() || checker.Encoding() != encoding.UTF_EBCDIC {
		t.Fatalf("Failed to handle UTF-EBCDIC LE file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.UTF_EBCDIC {
		t.Fatalf("Failed to handle UTF-EBCDIC LE file")
	}
}

func TestSCSU(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0x0E)
	checker.CheckNext(0xFE)
	checker.CheckNext(0xFF)
	if !checker.Validates() || checker.Encoding() != encoding.SCSU {
		t.Fatalf("Failed to handle SCSU file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.SCSU {
		t.Fatalf("Failed to handle SCSU file")
	}
}

func TestBOCU_1(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0xFB)
	checker.CheckNext(0xEE)
	checker.CheckNext(0x28)
	if !checker.Validates() || checker.Encoding() != encoding.BOCU_1 {
		t.Fatalf("Failed to handle BOCU-1 file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.BOCU_1 {
		t.Fatalf("Failed to handle BOCU-1 file")
	}
}

func TestGB_18030(t *testing.T) {
	checker := NewBomChecker()
	checker.CheckNext(0x84)
	checker.CheckNext(0x31)
	checker.CheckNext(0x95)
	checker.CheckNext(0x33)
	if !checker.Validates() || checker.Encoding() != encoding.GB_18030 {
		t.Fatalf("Failed to handle GB-18030 file")
	}
	checker.CheckNext(0x00)
	if !checker.Validates() || checker.Encoding() != encoding.GB_18030 {
		t.Fatalf("Failed to handle GB-18030 file")
	}
}
