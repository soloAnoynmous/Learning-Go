package main

import "fmt"

func main() {
	/*
		Working with slice operator
	*/

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	slice := arr[:]

	arr[2] = 23
	fmt.Println(arr)
	fmt.Println(slice)

	slice2 := slice[1:]  // start from index 1 to the end
	slice3 := slice[:2]  // start from index 0 to index 1
	slice4 := slice[1:2] // start from index 1 to index 1

	fmt.Println(slice2, slice3, slice4)

}
