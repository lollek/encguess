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

	printableAsciiChecker := NewPrintableAsciiChecker()
	asciiChecker := NewAsciiChecker()
	utf8Checker := NewUtf8Checker()
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
			printableAsciiChecker.CheckNext(data[i])
			asciiChecker.CheckNext(data[i])
			utf8Checker.CheckNext(data[i])
		}
	}

	switch {
	case printableAsciiChecker.Validates():
		fmt.Printf("%s\n", &printableAsciiChecker)
	case asciiChecker.Validates():
		fmt.Printf("%s\n", &asciiChecker)
	case utf8Checker.Validates():
		fmt.Printf("%s\n", &utf8Checker)
	default:
		println("Unknown encoding")
	}
}
