package main

import "fmt"

const (
	first  = iota // 0
	second = iota // 1
)

const (
	third  = iota + 6  // 0+6 = 6
	fourth = 2 << iota // 4
)

func main() {
	fmt.Println(first, second, third, fourth)
}
