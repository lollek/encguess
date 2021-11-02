package main

import "bytes"

func GuessEncodingFromBOM(data *[]byte, dataSize int) Encoding {
    if dataSize < 2 || data == nil {
        return UNKNOWN
    }

    var twoBytes = (*data)[:2]
    if bytes.Equal(twoBytes, []byte {0xFE, 0xFF}) {
        return UTF16_BE
    }

    if bytes.Equal(twoBytes, []byte {0xFF, 0xFE}) {
        if dataSize >= 4 && bytes.Equal((*data)[2:4], []byte{0x00, 0x00}) {
            return UTF32_LE
        } else {
            return UTF16_LE
        }
    }

    if dataSize < 3 {
        return UNKNOWN
    }
    var threeBytes = (*data)[:3]

    if bytes.Equal(threeBytes, []byte {0xF7, 0x64, 0x4C}) {
        return UTF1
    }

    if bytes.Equal(threeBytes, []byte {0x2B, 0x2F, 0x76}) {
        return UTF7
    }

    if bytes.Equal(threeBytes, []byte {0xEF, 0xBB, 0xBF}) {
        return UTF8
    }

    if dataSize < 4 {
        return UNKNOWN
    }

    var fourBytes = (*data)[:4]
    if bytes.Equal(fourBytes, []byte {0x00, 0x00, 0xFE, 0xFF}) {
        return UTF32_BE
    }

    return UNKNOWN
}

