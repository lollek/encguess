package main


import (
	"bytes"
	"github.com/lollek/encodingutil/encoding"
)

// Parses the data received and sees if there is any BOM which matches a known
// BOM - in that case the encoding from that BOM is returned. Else UNKNOWN is
// returned.
func GuessEncodingFromBOM(data *[]byte, dataSize int) encoding.Encoding {
	if dataSize < 2 || data == nil {
		return encoding.UNKNOWN
	}

	var twoBytes = (*data)[:2]
	if bytes.Equal(twoBytes, []byte{0xFE, 0xFF}) {
		return encoding.UTF16_BE
	}

	if bytes.Equal(twoBytes, []byte{0xFF, 0xFE}) {
		if dataSize >= 4 && bytes.Equal((*data)[2:4], []byte{0x00, 0x00}) {
			return encoding.UTF32_LE
		} else {
			return encoding.UTF16_LE
		}
	}

	if dataSize < 3 {
		return encoding.UNKNOWN
	}
	var threeBytes = (*data)[:3]

	if bytes.Equal(threeBytes, []byte{0xF7, 0x64, 0x4C}) {
		return encoding.UTF1
	}

	if bytes.Equal(threeBytes, []byte{0x2B, 0x2F, 0x76}) {
		return encoding.UTF7
	}

	if bytes.Equal(threeBytes, []byte{0xEF, 0xBB, 0xBF}) {
		return encoding.UTF8
	}

	if bytes.Equal(threeBytes, []byte{0x0E, 0xFE, 0xFF}) {
		return encoding.SCSU
	}

	if bytes.Equal(threeBytes, []byte{0xFB, 0xEE, 0x28}) {
		return encoding.BOCU_1
	}

	if dataSize < 4 {
		return encoding.UNKNOWN
	}

	var fourBytes = (*data)[:4]
	if bytes.Equal(fourBytes, []byte{0x00, 0x00, 0xFE, 0xFF}) {
		return encoding.UTF32_BE
	}

	if bytes.Equal(fourBytes, []byte{0xDD, 0x73, 0x66, 0x73}) {
		return encoding.UTF_EBCDIC
	}

	if bytes.Equal(fourBytes, []byte{0x84, 0x31, 0x95, 0x33}) {
		return encoding.GB_18030
	}

	return encoding.UNKNOWN
}
