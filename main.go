// package for executable build
package main

// formatting strings and printing messages
import "fmt"

// main function - only have one on your entrypoint and not on other files
func main() {

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

	// --------------------------------

	// PACKAGE: "strings"

	// EXAMPLE: fmt.Println(strings.Contains(helloWorld, "Hello"))
	// // => returns true
	// EXAMPLE: fmt.Println(strings.ReplaceAll(helloWorld, "Hello", "Hi"))
	// // => returns "Hi honey!"

}

// To run : `go run [file]`
