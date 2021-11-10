package checker

import (
	"github.com/lollek/encodingutil/encoding"
)

type BomChecker struct {
	validates          bool
	encodingFromBom    encoding.Encoding
	doneParsing        bool
	currentIndex       int
	potentialEncodings []encoding.Encoding
}

func NewBomChecker() *BomChecker {
	return &BomChecker{
		validates:       true,
		encodingFromBom: encoding.UNKNOWN,
		doneParsing:     false,
		currentIndex:    0,
		potentialEncodings: []encoding.Encoding{
			encoding.UTF1, encoding.UTF7, encoding.UTF8,
			encoding.UTF16_BE, encoding.UTF16_LE, encoding.UTF32_BE,
			encoding.UTF32_LE, encoding.UTF_EBCDIC, encoding.SCSU,
			encoding.BOCU_1, encoding.GB_18030,
		},
	}
}

func (checker *BomChecker) Encoding() encoding.Encoding {
	return checker.encodingFromBom
}

func (checker *BomChecker) Validates() bool {
	return checker.validates
}

func (checker *BomChecker) Probability() Probability {
	return HIGH
}

func (checker *BomChecker) CheckNext(character byte) {
	if checker.doneParsing {
		return
	}

	bomForEncoding := map[encoding.Encoding][]byte{
		encoding.UTF1:       []byte{0xF7, 0x64, 0x4C},
		encoding.UTF7:       []byte{0x2B, 0x2F, 0x76},
		encoding.UTF8:       []byte{0xEF, 0xBB, 0xBF},
		encoding.UTF16_BE:   []byte{0xFE, 0xFF},
		encoding.UTF16_LE:   []byte{0xFF, 0xFE},
		encoding.UTF32_BE:   []byte{0x00, 0x00, 0xFE, 0xFF},
		encoding.UTF32_LE:   []byte{0xFF, 0xFE, 0x00, 0x00},
		encoding.UTF_EBCDIC: []byte{0xDD, 0x73, 0x66, 0x73},
		encoding.SCSU:       []byte{0x0E, 0xFE, 0xFF},
		encoding.BOCU_1:     []byte{0xFB, 0xEE, 0x28},
		encoding.GB_18030:   []byte{0x84, 0x31, 0x95, 0x33},
	}

	numberOfStillParsingEncodings := 0
	for i := 0; i < len(checker.potentialEncodings); {
		enc := checker.potentialEncodings[i]
		bom := bomForEncoding[enc]

		if checker.currentIndex+1 >= len(bom) {
			checker.encodingFromBom = enc
			i++
		} else if bom[checker.currentIndex] == character {
			numberOfStillParsingEncodings++
			i++
		} else {
			checker.potentialEncodings = append(checker.potentialEncodings[:i], checker.potentialEncodings[i+1:]...)
		}
	}

	checker.currentIndex++

	if numberOfStillParsingEncodings == 0 {
		checker.doneParsing = true

		switch len(checker.potentialEncodings) {
		case 0:
			checker.validates = false
		case 1:
			checker.encodingFromBom = checker.potentialEncodings[0]
		case 2:
			longestBom := encoding.UNKNOWN
			longestBomLength := 0
			for _, enc := range checker.potentialEncodings {
				bomLength := len(bomForEncoding[enc])
				if bomLength > longestBomLength {
					longestBom = enc
					longestBomLength = bomLength
				}
			}
			checker.encodingFromBom = longestBom
		}
	}
}
