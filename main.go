package main

import "fmt"

func main(){
// Some functions accepts a variable number of parameters
// This kind of params are called variadic parameters and are indicated with this notation in documentation: ...<some type of data>, for example ...<int>
//In case we want to accept any kind of data, the notation is ...{interface}, since in Go any variable is also an interface
	fmt.Println("Hi everybody", 27, true);
//n, err := fmt.Println(...{interface}) where n is the number of printed char
	n, err :=fmt.Println("Test");
	if err==nil{
		fmt.Printf("I have printed %d characters\n", n)
	}
	m, _ := fmt.Println("Now I have printed:")
	fmt.Printf("%d characters\n", m)
}
