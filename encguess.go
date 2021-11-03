package main

import (
	"fmt"
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
		fmt.Printf("Failed to open file `%s`: %s\n", fileName, err)
		os.Exit(1)
	}

	data := make([]byte, os.Getpagesize())
	count, err := file.Read(data)
	if err != nil {
		fmt.Printf("Failed to read from file `%s`: %s\n", fileName, err)
		os.Exit(1)
	}

	if encoding := GuessEncodingFromBOM(&data, count); encoding != UNKNOWN {
		fmt.Printf("I'm guessing the encoding is %s because of the BOM\n", encoding)
		os.Exit(0)
	}

	fmt.Printf("Failed to guess the encoding :(\n")
}
