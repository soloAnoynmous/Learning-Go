package main

import "fmt"

func main() {
	/*
		Creating arrays : way1
	*/

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	/*
		Creating arrays : way2
	*/
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

}
