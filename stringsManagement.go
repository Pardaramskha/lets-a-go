package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello %v \n", name)
}

func checkAge(name string) float64 {

	// for maps like this, don't forget to add a colon even on the last element
	charactersAge := map[string]float64{
		"Keira":        21,
		"Rapto":        18,
		"Raeza":        33,
		"Charlotte":    28,
		"Pardaramskha": 9715,
	}

	// we can update maps at any moment with, for example
	// -- charactersAge["Keira"] = 22

	var _age float64

	for key, value := range charactersAge {
		if key == name {
			_age = value
			break
		}
	}

	// This only checks directly without handling different cases
	//return charactersAge[name]

	return _age
}

func updateAge(age *int) {
	*age++
}

func pointersAndMemory() {
	fmt.Println("---- POINTERS AND MEMORY")
	keirasAge := 21

	// when passed to a function, some data can be updated to the original without returning a new value

	// NON POINTER VALUES: strings, ints, floats, booleans, arrays, structs
	// // Originals WON'T be updated because Go creates a copy of the value
	// POINTERS VALUES: maps, slices, functions
	// // Originals WILL be updated

	// To change originals of non pointer values, we have to create a pointer first then get its value
	// Pointers are stocked into another memory block that points towards the original value
	/*
			|--name---|--namePointer---|
			|  0x001  |  0x002         |
		    |---------|----------------|
		    |  "azi"  |  p0x001        |
		    |---------|----------------|
	*/
	// To create pointer, just add "&" before the variable name. Logging this will show memory address
	// name := "azi"
	// namePointer := &name --- logs some math gibberish
	// To get a value associated to a pointer, add "*" before the pointer name
	// _name := *namePointer --- logs "azi"

	_keirasAge := &keirasAge
	fmt.Println("Memory address of variable is:", _keirasAge)
	fmt.Println("Value at this memory address is:", *_keirasAge)

	updateAge(_keirasAge)

	fmt.Println("Now updating initial pointed value", keirasAge)
}

func structsAndObjects() {

}
