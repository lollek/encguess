package main

func GuessEncodingFromBOM(data *[]byte, dataSize int) Encoding {
    if dataSize < 2 || data == nil {
        return UNKNOWN
    }

    if (*data)[0] == 0xFE && (*data)[1] == 0xFF {
        return UTF16_BE
    }

    if (*data)[0] == 0xFF && (*data)[1] == 0xFE {
        if dataSize >= 4 && (*data)[2] == 0x00 && (*data)[3] == 0x00 {
            return UTF32_LE
        } else {
            return UTF16_LE
        }
    }

    if dataSize < 3 {
        return UNKNOWN
    }

    if (*data)[0] == 0xF7 && (*data)[1] == 0x64 && (*data)[2] == 0x4C {
        return UTF1
    }

    if (*data)[0] == 0x2B && (*data)[1] == 0x2F && (*data)[2] == 0x76 {
        return UTF7
    }

    if (*data)[0] == 0xEF && (*data)[1] == 0xBB && (*data)[2] == 0xBF {
        return UTF8
    }

    if dataSize < 4 {
        return UNKNOWN
    }

    if (*data)[0] == 0x00 && (*data)[1] == 0x00 && (*data)[2] == 0xFE && (*data)[3] == 0xFF {
        return UTF32_BE
    }

    return UNKNOWN
}

