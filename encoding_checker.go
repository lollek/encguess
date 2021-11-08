package main

type EncodingChecker interface {
	Validates() bool
	CheckNext(character byte)
}
