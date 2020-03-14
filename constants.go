package main

import "fmt"

func main() {

	// without type declaration
	const pi = 3.14
	fmt.Println(pi)

	// with type declaration
	const p int = 3
	fmt.Println(p)

	// casting : int cannot be truncated to float
	fmt.Println(float32(p) + 3.14)

}
