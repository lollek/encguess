package fileparser

import (
	"fmt"
	"github.com/lollek/encodingutil/checker"
	"github.com/lollek/encodingutil/encoding"
	"io"
	"os"
)

type PotentialEncodingPair struct {
	Encoding    encoding.Encoding
	Probability checker.Probability
}

type ParseFileResult struct {
	PotentialEncodings []PotentialEncodingPair
}

func ParseFile(fileName string) (*ParseFileResult, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file `%s`: %s\n", fileName, err)
	}

	data := make([]byte, os.Getpagesize())

	checkers := []checker.EncodingChecker{
		checker.NewBomChecker(),
		checker.NewPrintableAsciiChecker(),
		checker.NewAsciiChecker(),
		checker.NewUtf8Checker(),
		checker.NewIso88591Checker(),
	}
loop:
	for {
		count, err := file.Read(data)
		switch err {
		case io.EOF:
			break loop
		case nil:
			break
		default:
			return nil, fmt.Errorf("Error reading from file `%s`: %s\n", fileName, err)
		}

		for i := 0; i < count; i++ {
			character := data[i]
			for checkerIndex := 0; checkerIndex < len(checkers); {
				checker := checkers[checkerIndex]
				checker.CheckNext(character)
				if !checker.Validates() {
					checkers = append(checkers[:checkerIndex], checkers[checkerIndex+1:]...)
				} else {
					checkerIndex++
				}
			}
			if len(checkers) == 0 {
				break loop
			}
		}
	}

	potentialEncodings := make([]PotentialEncodingPair, len(checkers))
	for i, checker := range checkers {
		potentialEncodings[i] = PotentialEncodingPair{
			Encoding:    checker.Encoding(),
			Probability: checker.Probability(),
		}
	}

	return &ParseFileResult{
		PotentialEncodings: potentialEncodings,
	}, nil
}
