// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func printGlobalVal() {
	fmt.Printf("%T, local var = %v\n", val, val)
}

func setVal(value string) {
	val = value
}

func main() {
	// create a local variable "val" and assign it an int value
	val := 100
	// print the value of the local variable "val"

	fmt.Printf("%T, local var = %v\n", val, val)

	// print the value of the package-level variable "val"
	printGlobalVal()

	// update the package-level variable "val" and print it again
	setVal("updated global")
	printGlobalVal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local var = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 300
	fmt.Printf("local var = %v\n", val)
}
