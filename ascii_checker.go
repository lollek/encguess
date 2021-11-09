package main

type AsciiChecker struct {
	validates bool
}

func NewAsciiChecker() *AsciiChecker {
	return &AsciiChecker{
		validates: true,
	}
}

func (checker *AsciiChecker) Encoding() Encoding {
	return ASCII
}

func (checker *AsciiChecker) Validates() bool {
	return checker.validates
}

func (checker *AsciiChecker) CheckNext(character byte) {
	checker.validates = checker.validates && character&0x80 == 0
}
