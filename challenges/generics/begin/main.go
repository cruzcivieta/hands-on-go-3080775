// challenges/generics/begin/main.go
package main

import "fmt"

// Part 1: print function refactoring
func printAny[T any](value T) {
	fmt.Println(value)
}

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
type numeric interface {
	~int | ~float64
}

func sumAny[T numeric](values ...T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println("result", sumAny(2.5, 2.7, 3.9))
	fmt.Println("result", sumAny(2, 3, 5))
}
