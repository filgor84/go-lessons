package main

import "fmt"

//Direct declaration (x:=42) is cool, but it doesn't work outside the scope of a function

var x = 42

//The keyword var should be used in this case

var z int

//assignment without specifying a value
//a default value is assigned to z (0)
//0 for all integer types,
//0.0 for floating point numbers,
//false for booleans,
//"" for strings,
//nil for interfaces, slices, channels, maps, pointers and functions.

func main() {
	//Direct declaration
	y := 43
	fmt.Println(x)
	fmt.Println(y)
	foo()

}

func foo() {
	fmt.Println(z)

}
