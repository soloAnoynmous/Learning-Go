package main

import "fmt"

func main() {
	m := map[string]int{"foo": 42} // map<string,int>
	fmt.Println(m)
	fmt.Println(m["foo"])

	//delete
	delete(m, "foo")

	fmt.Println(m)
}
