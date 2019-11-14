package main

import "fmt"

func main() {
	fmt.Println("Hello everybody!!!")
	foo()
	fmt.Println("Even numbers to 100")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar. \n Time to exit!")
}
