package main

import "fmt"

func main() {

	/*
		Pointers can be used in two ways.
	*/

	// Way 1
	// creating a pointer to a string
	var firstName *string = new(string)
	*firstName = "satyam"
	fmt.Println(firstName, *firstName)

	// Way 2 - using addressOf operator
	lastName := "satyam"
	ptr := &lastName
	fmt.Println(ptr, *ptr)
}
