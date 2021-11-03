package main

func IsPrintableASCII(character byte) bool {
	if 0x20 <= character && character <= 0x7E {
		return true
	}

	switch character {
	case 0x0A /* \n */, 0x0B /* \v */, 0x0D /* \r */ :
		return true
	default:
		return false
	}
}

func IsASCII(character byte) bool {
	return character&0x80 == 0
}
