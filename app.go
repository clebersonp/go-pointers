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

	adultYears := getAdultYears(&age) // accessing the memory address and pass it as pointer
	fmt.Println(adultYears)

	fmt.Println("-----------------------")
	fmt.Println("Example using pointer to mutate values directly from pointers")
	fmt.Println("Age:", age)
	editAgeToAdultYears(&age)
	fmt.Println(age)
}

// age *int means that will pass a pointer (copy of address memory) not copy value
func getAdultYears(age *int) int {
	return *age - 18 //*age means to access the real data value from memory address
}

func editAgeToAdultYears(age *int) {
	*age -= 18 // accessing the value from pointer and change it directly
}
