package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // *int means a pointer of int type
	agePointer = &age   // accessing the age memory address (pointer)
	fmt.Println("Age pointer:", agePointer)

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
