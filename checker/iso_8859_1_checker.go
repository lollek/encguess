package checker

import (
	"github.com/lollek/encodingutil/encoding"
)

type Iso88591Checker struct {
	validates bool
}

func NewIso88591Checker() *Iso88591Checker {
	return &Iso88591Checker{
		validates: true,
	}
}

func (checker *Iso88591Checker) Encoding() encoding.Encoding {
	return encoding.ISO_8859_1
}

func (checker *Iso88591Checker) Validates() bool {
	return checker.validates
}

func (checker *Iso88591Checker) Probability() Probability {
	return MEDIUM
}

func (checker *Iso88591Checker) CheckNext(character byte) {
	if !checker.Validates() {
		return
	}

	if 0x80 <= character && character < 0xA0 {
		checker.validates = false
	}
}
