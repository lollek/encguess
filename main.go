package main

import (
	"flag"
	"github.com/lollek/encodingutil/fileparser"
	"fmt"
	"os"
	"sort"
)

var (
	printFileName *bool
	printAllPotentialEncodings *bool
)

func printParseFileResult(fileName string, result *fileparser.ParseFileResult) {
	results := result.PotentialEncodings
	sort.Slice(results, func(a, b int) bool {
		return results[a].Probability > results[b].Probability
	})

	if *printFileName {
		fmt.Printf("%s: ", fileName)
	}

	numResults := len(results)
	if numResults > 1 && !*printAllPotentialEncodings {
		numResults = 1
	}
	switch numResults {
	case 0:
		fmt.Printf("Unknown\n")
	case 1:
		fmt.Printf("%s\n", results[0].Encoding)
	default:
		for i, guess := range results {
			if i != 0 {
				fmt.Printf(" or ")
			}
			fmt.Printf("%s", guess.Encoding)
		}
		println()
	}
}

func main() {
	printFileName = flag.Bool("f", false, "Prefix output with the file name")
	printAllPotentialEncodings = flag.Bool("a", false, "Print all potential encodings, instead of just the most likely. Will be in order of likelyness")
	flag.Parse()

	for _, fileName := range flag.Args() {
		result, err := fileparser.ParseFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		} else {
			printParseFileResult(fileName, result)
		}
	}
}
