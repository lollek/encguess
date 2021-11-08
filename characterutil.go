package main

func CreatePrintableAsciiChecker() func(byte) bool {
	isPrintableAscii := true

	return func(character byte) bool {
		if !isPrintableAscii {
			return false
		}
		if 0x20 <= character && character <= 0x7E {
			return true
		}

		switch character {
		case 0x0A /* \n */, 0x0B /* \v */, 0x0D /* \r */ :
			return true
		default:
			isPrintableAscii = false
			return isPrintableAscii
		}
	}
}

func CreateAsciiChecker() func(byte) bool {
	isAscii := true

	return func(character byte) bool {
		if !isAscii {
			return false
		}
		isAscii = character&0x80 == 0
		return isAscii
	}
}

func CreateUtf8Checker() func(byte) bool {
	// UTF-8 characters can be one to four bytes, depending on the starting
	// flag. See https://en.wikipedia.org/wiki/UTF-8.

	isUTF8 := true

	// Current index is index of the UTF-8 character.
	currentIndex := 0

	// Max index is how many characters there should be in the character.
	maxIndex := 0

	return func(character byte) bool {
		if !isUTF8 {
			return false
		}

		if currentIndex == 0 {
			if character&0x80 == 0 {
				currentIndex = 0
				maxIndex = 0
				return true
			}
			currentIndex = 1
			if character&0xE0 == 0xC0 {
				maxIndex = 1
				return true
			}
			if character&0xF0 == 0xE0 {
				maxIndex = 2
				return true
			}
			if character&0xF8 == 0xF0 {
				maxIndex = 3
				return true
			}
			isUTF8 = false
			return isUTF8
		}

		if character&0xC0 == 0x80 {
			if currentIndex == maxIndex {
				currentIndex = 0
			} else {
				currentIndex += 1
			}
			return true
		}
		isUTF8 = false
		return isUTF8
	}
}
