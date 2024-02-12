package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // *int means a pointer of int type
	agePointer = &age   // accessing the age memory address (pointer)
	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age pointer real value:", *agePointer) // *pointer to get the real value from pointer(memory address)
	// For a pointer, the 'null value' (default value) is nil

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
