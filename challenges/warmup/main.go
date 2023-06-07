package main

import "fmt"

var name string = "John"
var age uint8 = 32
var template string = "%v is years %v years old\n"

func getPersonalInfo() (string, uint8) {
	return name, age
}

func (person person) printPersonalInfo() {
	fmt.Printf(template, person.name, person.age)
}

func (person person) setAgeCopy(age uint8) {
	person.age = age
}

func (person *person) setAgePointer(age uint8) {
	person.age = age
}

type person struct {
	name string
	age  uint8
}

func main() {
	// 1
	fmt.Println("Hello, word")

	// 2
	fmt.Printf(template, name, age)

	// 2.1
	name, age := getPersonalInfo()
	fmt.Printf(template, name, age)

	// 2.1
	fmt.Println(name)

	// 3
	john := person{
		name: "John",
		age:  32,
	}
	fmt.Printf(template, john.name, john.age)

	// 3.1
	john.printPersonalInfo()

	// 3.2
	john.setAgeCopy(33)
	john.printPersonalInfo()

	john.setAgePointer(34)
	john.printPersonalInfo()

	// 4
	x := 40
	fmt.Printf("memory addres %v\n", &x)

	// 4.1
	y := &x
	fmt.Printf("value of y is %v\n", *y)

	// 5
	week := [](string){"Monday", "Tuesday", "Wensday", "Thursday", "Friday"}
	for _, day := range week {
		fmt.Printf("Current day is %v\n", day)
	}

	// 5.1
	week = append(week, "Saturday", "Sunday")
	for index, day := range week {
		if index > 5 {
			fmt.Println("It's weekend")
		} else {
			fmt.Printf("Current day is %v\n", day)
		}
	}
}
