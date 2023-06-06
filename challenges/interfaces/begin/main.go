// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }

func (letterCounter letterCounter) name() string {
	return letterCounter.identifier
}

func (letterCounter letterCounter) count(input string) int {
	var counter int = 0
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			counter++
		}
	}
	return counter
}

type numberCounter struct{ designation string }

func (numberCounter numberCounter) name() string {
	return numberCounter.designation
}

func (numberCounter numberCounter) count(input string) int {
	var counter int = 0
	for _, letter := range input {
		if unicode.IsDigit(letter) {
			counter++
		}
	}
	return counter
}

type symbolCounter struct{ label string }

func (symbolCounter symbolCounter) name() string {
	return symbolCounter.label
}

func (symbolCounter symbolCounter) count(input string) int {
	var counter int = 0
	for _, letter := range input {
		if !unicode.IsDigit(letter) && !unicode.IsLetter(letter) {
			counter++
		}
	}
	return counter
}

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	spew.Dump(data)

	// call doAnalysis and pass in the data and the counters
	letterCounter := letterCounter{identifier: "letterCounter"}
	numberCounter := numberCounter{designation: "numberCounter"}
	symbolCounter := symbolCounter{label: "symbolCounter"}

	analysis := doAnalysis(data, letterCounter, numberCounter, symbolCounter)

	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
