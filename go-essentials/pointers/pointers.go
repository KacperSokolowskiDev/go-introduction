package main

import "fmt"

func main() {
	age := 24 //regular variable

	// vat agePointer *int = &age
	agePointer := &age

	fmt.Println("Age:", *agePointer) //dereferencing agePointer

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	//placing result in dereferenced pointer, alternative to returning
	*age = *age - 18 //overwritting result under age address
}