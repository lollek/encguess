package checker

import (
	"github.com/lollek/encodingutil/encoding"
)

type EncodingChecker interface {
	Validates() bool
	Probability() Probability
	Encoding() encoding.Encoding
	CheckNext(character byte)
}
