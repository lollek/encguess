package main

import "testing"

func TestUTF1(t *testing.T) {
    var data = []byte {0xF7, 0x64, 0x4C}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF1 {
        t.Fatalf("Failed to handle UTF-1 file. Found %s", result)
    }
}

func TestUTF7(t *testing.T) {
    var data = []byte {0x2B, 0x2F, 0x76}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF7 {
        t.Fatalf("Failed to handle UTF-7 file. Found %s", result)
    }
}

func TestUTF8(t *testing.T) {
    var data = []byte {0xEF, 0xBB, 0xBF}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF8 {
        t.Fatalf("Failed to handle UTF-8 file. Found %s", result)
    }
}

func TestUTF16_BE(t *testing.T) {
    var data = []byte {0xFE, 0xFF}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF16_BE {
        t.Fatalf("Failed to handle UTF-16 BE file. Found %s", result)
    }
}

func TestUTF16_LE(t *testing.T) {
    var data = []byte {0xFF, 0xFE}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF16_LE {
        t.Fatalf("Failed to handle UTF-16 LE file. Found %s", result)
    }
}

func TestUTF32_BE(t *testing.T) {
    var data = []byte {0x00, 0x00, 0xFE, 0xFF}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF32_BE {
        t.Fatalf("Failed to handle UTF-32 BE file. Found %s", result)
    }
}

func TestUTF32_LE(t *testing.T) {
    var data = []byte {0xFF, 0xFE, 0x00, 0x00}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF32_LE {
        t.Fatalf("Failed to handle UTF-32 LE file. Found %s", result)
    }
}

func TestUTF_EBCDIC(t *testing.T) {
    var data = []byte {0xDD, 0x73, 0x66, 0x73}
    if result := GuessEncodingFromBOM(&data, len(data)); result != UTF_EBCDIC {
        t.Fatalf("Failed to handle UTF-EBCDIC file. Found %s", result)
    }
}

func TestSCSU(t *testing.T) {
    var data = []byte {0x0E, 0xFE, 0xFF}
    if result := GuessEncodingFromBOM(&data, len(data)); result != SCSU {
        t.Fatalf("Failed to handle SCSU file. Found %s", result)
    }
}

func TestBOCU_1(t *testing.T) {
    var data = []byte {0xFB, 0xEE, 0x28}
    if result := GuessEncodingFromBOM(&data, len(data)); result != BOCU_1 {
        t.Fatalf("Failed to handle BOCU_1 file. Found %s", result)
    }
}

func TestGB_18030(t *testing.T) {
    var data = []byte {0x84, 0x31, 0x95, 0x33}
    if result := GuessEncodingFromBOM(&data, len(data)); result != GB_18030 {
        t.Fatalf("Failed to handle GB_18030 file. Found %s", result)
    }
}

func TestNullFile(t *testing.T) {
    if result := GuessEncodingFromBOM(nil, 0); result != UNKNOWN {
        t.Fatalf("Failed to handle nil file. Expected UNKNOWN. Found %s", result)
    }
    if result := GuessEncodingFromBOM(nil, 1); result != UNKNOWN {
        t.Fatalf("Failed to handle nil file. Expected UNKNOWN. Found %s", result)
    }
}

func TestFileLengths(t *testing.T) {
    for size := 0; size < 4096; size++ {
        var data = make([]byte, size)
        if result := GuessEncodingFromBOM(&data, size); result != UNKNOWN {
            t.Fatalf("Failed to handle size %d file. Expected UNKNOWN. Found %s", size, result)
        }
    }
}
