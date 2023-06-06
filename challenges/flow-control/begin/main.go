// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from error: ", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	filename := os.Args[1]
	if filename == "" {
		panic("Not provided any filename")
	}
	data, fileError := os.ReadFile(filename)
	if fileError != nil {
		panic("Error reading the provided filename")
	}

	// convert the bytes to a string
	text := string(data[:])
	spew.Dump(text)

	// initialize a map to store the counts
	counts := make(map[string]int, 0)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(text)

	// capture the length of the words slice
	wordsLenght := len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	counts["words"] = wordsLenght
	for _, word := range words {
		for _, letter := range word {
			isDigit := unicode.IsDigit(letter)
			isLetter := unicode.IsLetter(letter)
			switch {
			case isDigit:
				counts["numbers"] += 1
			case isLetter:
				counts["letters"] += 1
			default:
				counts["symbols"] += 1
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
