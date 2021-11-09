package main

type Utf8Checker struct {
	validates    bool
	currentIndex int
	maxIndex     int
}

func NewUtf8Checker() *Utf8Checker {
	return &Utf8Checker{
		validates:    true,
		currentIndex: 0,
		maxIndex:     0,
	}
}

func (checker *Utf8Checker) Encoding() Encoding {
	return UTF8
}

func (checker *Utf8Checker) Validates() bool {
	return checker.validates
}

func (checker *Utf8Checker) CheckNext(character byte) {
	// UTF-8 characters can be one to four bytes, depending on the starting
	// flag. See https://en.wikipedia.org/wiki/UTF-8.

	if !checker.Validates() {
		return
	}

	if checker.currentIndex == 0 {
		if character&0x80 == 0 {
			checker.currentIndex = 0
			checker.maxIndex = 0
		} else if character&0xE0 == 0xC0 {
			checker.currentIndex = 1
			checker.maxIndex = 1
		} else if character&0xF0 == 0xE0 {
			checker.currentIndex = 1
			checker.maxIndex = 2
		} else if character&0xF8 == 0xF0 {
			checker.currentIndex = 1
			checker.maxIndex = 3
		} else {
			checker.validates = false
		}
	} else if character&0xC0 == 0x80 {
		if checker.currentIndex == checker.maxIndex {
			checker.currentIndex = 0
		} else {
			checker.currentIndex += 1
		}
	} else {
		checker.validates = false
	}
}
