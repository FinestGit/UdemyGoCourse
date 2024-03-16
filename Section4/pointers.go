package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age:", *agePointer)
	fmt.Println("Adult years:", getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
