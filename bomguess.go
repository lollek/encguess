package main

func GuessEncodingFromBOM(data *[]byte, dataSize int) Encoding {
    if dataSize == 0 || data == nil {
        return UNKNOWN
    }

    if (*data)[0] == 0xEF && (*data)[1] == 0xBB && (*data)[2] == 0xBF {
        return CHARSET_UTF8
    }

    return UNKNOWN
}

