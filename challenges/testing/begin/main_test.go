// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	testData := []struct {
		name  string
		input string
		want  int
	}{
		{"no text", "", 0},
		{"no letters", "/#00", 0},
		{"with a letter", "/a2_32_^_&/", 1},
		{"more letters", "/812:%6ab/", 2},
	}
	letterCounter := letterCounter{identifier: "identifier"}
	for _, value := range testData {
		result := letterCounter.count(value.input)

		if result != value.want {
			t.Errorf("got %v want %v", result, value.want)
		}
	}
}

// write a test for numberCounter.count
func TestNumberLetterCount(t *testing.T) {
	testData := []struct {
		name  string
		input string
		want  int
	}{
		{"no text", "", 0},
		{"no numbers", "/#", 0},
		{"with a number", "/abc_1,?/", 1},
		{"more numbers", "abc_12&8_^", 3},
	}
	numberCounter := numberCounter{designation: "number"}
	for _, value := range testData {
		result := numberCounter.count(value.input)

		if result != value.want {
			t.Errorf("got %v want %v", result, value.want)
		}
	}
}

// write a test for symbolCounter.count
func TestSymbolLetterCount(t *testing.T) {
	testData := []struct {
		name  string
		input string
		want  int
	}{
		{"no text", "", 0},
		{"no symbol", "00aa", 0},
		{"with symbols", "abc_1,?/", 4},
		{"more symbols", "abc_12&8_^_", 5},
	}
	symbolCounter := symbolCounter{label: "number"}
	for _, value := range testData {
		result := symbolCounter.count(value.input)

		if result != value.want {
			t.Errorf("got %v want %v", result, value.want)
		}
	}
}
