// package for executable build
package main

// fmt: formatting strings and printing messages
// sort: lots of methods to sort stuff
import (
	"fmt"
	"sort"
	"strings"
)

// main function - only have one on your entrypoint and not on other files OF THE SAME PACKAGE
func main() {
	//sayHello("Douchebag")

	//essentials()
	//loops()

	//mappingAndObjects()
	//pointersAndMemory()

	keira := newCharacter("Keira")
	fmt.Println(keira.stuffFormat())
}

func essentials() {
	// GO vars will automatically infer type, but specify it anyway, it'll help
	// watch out : we NEED to use variables otherwise the compiler goes sob sob

	// STRING DOCUMENTATION
	// string only accept double quotes
	var helloWorld string = "Hello honey!"
	// shorthand variables can only be used INSIDE functions
	shorthand := "Those are the first variables written - good luck!"

	// PRINT DOCUMENTATION
	// .Println creates a new line
	// .Print prints on the same line unless \n is added
	// .Printf formats the line (does not add a new line)
	// // %_ = format specifier (v = variable, q = wraps string between quotes, T = prints type, f = full float, [decimal number]f = fixed number of decimals)
	// .Sprintf creates formatted message and saves it (must be used as variable value)
	fmt.Println(helloWorld, shorthand)
	fmt.Printf("The first variable was %v \n", helloWorld)

	// INT DOCUMENTATION
	// differences can be made between int, int8, int16, int32, int64, uint etc.
	// this specifies the bits for an int - they have a specific range of numbers that throw an error if not respected
	// the particular uint (with 8/16/32.. variables) only accepts positive numbers

	// FLOAT DOCUMENTATION
	// we HAVE to specify the number of bits with type float32, float64
	// defaults to float64 while shorthanded

	// ARRAY DOCUMENTATION
	// var array [number of items in array]type of the items = [idem]idem{array content}
	// shorthand: var array = [idem]idem{array content}
	// WARNING: Array length cant' change!
	array := [3]string{"keira", "horkos", "pardaramskha"}
	// the len function prints array length
	fmt.Println(array, len(array))

	// SLICES DOCUMENTATION
	// The arrays as we know 'em - uses arrays under the hood
	// var slice = []int{val1, val2, val3}
	// change slices/arrays: slice[index] = newValue
	// append item to slice: slice = append(slice, val4)
	// get range of elements from slice: range := slice[startVal:endVal]
	// EXAMPLE:
	// // mySlice := []int{10,20,30}
	// // myRange := mySlice[1:2] (returns 20,30)
	// // myRange := mySlice[:1] (returns 10,20 - it's UP TO position X)
	// // myRange := mySlice[1:] (returns 20,30 - it's FROM position X)

	charactersSlice := []string{"rapto", "keira", "horkos", "pardaramskha", "azi"}

	// PACKAGE: "sort"

	// the sort methods will modify the original slice - no need to register that in a new variable for it to take effect
	sort.Strings(charactersSlice)

	// the sort.Search___ will return the index of desired element, with args (slice to search, element to find)
	fmt.Println(charactersSlice, sort.SearchStrings(charactersSlice, "keira"))

	// --------------------------------

	// PACKAGE: "strings"

	// EXAMPLE: fmt.Println(strings.Contains(helloWorld, "Hello"))
	// // => returns true
	// EXAMPLE: fmt.Println(strings.ReplaceAll(helloWorld, "Hello", "Hi"))
	// // => returns "Hi honey!"
}

func loops() {

	charactersSlice := []string{"rapto", "keira", "horkos", "pardaramskha", "azi"}

	fmt.Println("---- SIMPLE LOOPS")
	// simple loop with initialisation, condition and increment per loop
	//for x := 0; x < len(charactersSlice); x++ {
	//	fmt.Println(charactersSlice[x])
	//}

	fmt.Println("---- FOR/IN LOOPS")
	// "for/in" loop

	// example 1
	//for index, value := range charactersSlice {
	//	fmt.Printf("%v is number %v \n", value, index)
	//}

	// watch out for it returns specific things: index first, value second. Use underscore to not use a designed value

	// example 2
	//for _, name := range charactersSlice {
	//	fmt.Println(name)
	//}

	// example 3 with conditions and keywords
	for _, name := range charactersSlice {
		if name == "keira" {
			fmt.Printf("Main character: %v \n", toUppercase(name))
			// the continue keyword stops the code here but allows the loop to continue
			continue
		} else if name == "pardaramskha" {
			fmt.Println("We stop here, we don't need the chicken")
			// the break keyword just stops everything
			break
		}

		fmt.Printf("Character: %v \n", toUppercase(name))
	}

	// NOTE: altering the value in a loop DOES NOT change original slice values
}

func cycleNames(slice []string, function func(string)) {
	for _, value := range slice {
		function(value)
	}
}

// we have to specify the type of value returned before the braces
// to return multiple values, add parenthesis. Example below returns two strings
// -- func functionName(arg string) (string, string) { ... }
func toUppercase(name string) string {
	return strings.ToUpper(name)
}

func mappingAndObjects() {
	//age := checkAge("Idiot")
	//fmt.Println(age)
}

// To run : `go run [files]`
