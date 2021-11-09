package main

type EncodingChecker interface {
	Validates() bool
	Encoding() Encoding
	CheckNext(character byte)
}
