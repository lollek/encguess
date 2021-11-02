package main

import "testing"

func TestUTF8BOM(t *testing.T) {
    var data = []byte {0xef, 0xbb, 0xbf}
    if result := GuessEncodingFromBOM(&data, len(data)); result != CHARSET_UTF8 {
        t.Fatalf("Failed to handle UTF-8 file. Expected UTF-8. Found %s", result)
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
