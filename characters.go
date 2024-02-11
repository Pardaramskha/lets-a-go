package main

import "fmt"

type character struct {
	name  string
	level int
	items map[string]int
}

// create character
func newCharacter(name string) character {
	char := character{
		name:  name,
		level: 0,
		items: map[string]int{
			"basic stuff": 50,
			"knife":       25,
		},
	}

	return char
}

// level up character
// this function is a receiver associated with a particular custom struct
// so we can call it by typing instanceOfStruct.levelUp()
func (char character) stuffFormat() string {
	formattedString := "Character's stuff: \n"
	var totalPrice int = 0

	// list items

	// NOTE: the %-25v is used to format by adding space
	for key, value := range char.items {
		formattedString += fmt.Sprintf("%-25v ...%vT. \n", key+":", value)
		totalPrice += value
	}

	// total
	formattedString += fmt.Sprintf("%-25v ...%vT.", "Total:", totalPrice)

	return formattedString

}
