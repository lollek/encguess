package checker

import (
	"github.com/lollek/encodingutil/encoding"
)

type PrintableAsciiChecker struct {
	validates bool
}

func NewPrintableAsciiChecker() *PrintableAsciiChecker {
	return &PrintableAsciiChecker{
		validates: true,
	}
}

func (checker *PrintableAsciiChecker) Encoding() encoding.Encoding {
	return encoding.PRINTABLE_ASCII
}

func (checker *PrintableAsciiChecker) Validates() bool {
	return checker.validates
}

func (checker *PrintableAsciiChecker) CheckNext(character byte) {
	if !checker.Validates() {
		return
	}
	if 0x20 <= character && character <= 0x7E {
		return
	}

	switch character {
	case 0x0A /* \n */, 0x0B /* \v */, 0x0D /* \r */ :
		return
	default:
		checker.validates = false
	}
}
