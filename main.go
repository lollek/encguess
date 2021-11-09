package main

import (
	"fmt"
	"io"
	"os"
)

func printUsage(programName string) {
	fmt.Fprintf(os.Stderr, "%s: FILENAME\n", programName)
}

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file `%s`: %s\n", fileName, err)
		os.Exit(1)
	}

	data := make([]byte, os.Getpagesize())
	firstRun := true

	checkers := []EncodingChecker{
		NewPrintableAsciiChecker(),
		NewAsciiChecker(),
		NewUtf8Checker(),
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
			fmt.Fprintf(os.Stderr, "Error reading from file `%s`: %s\n", fileName, err)
			os.Exit(1)
		}

		if firstRun {
			firstRun = false
			if encoding := GuessEncodingFromBOM(&data, count); encoding != UNKNOWN {
				fmt.Printf("BOM-encoding: %s\n", encoding)
				os.Exit(0)
			}
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

	switch len(checkers) {
	case 0:
		println("Unknown")
	case 1:
		fmt.Printf("%s\n", checkers[0].Encoding())
	default:
		for checkerIndex := 0; checkerIndex < len(checkers); checkerIndex++ {
			fmt.Printf("%s\n", checkers[checkerIndex].Encoding())
		}
	}
}
