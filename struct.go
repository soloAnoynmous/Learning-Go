package main

import "fmt"

func main() {

	// Way 1
	type user struct {
		ID        int
		FirstName string
	}

	var u user
	u.ID = 1
	u.FirstName = "Satyam"
	fmt.Println(u)

	// Way 2
	u2 := user{ID: 1,
		FirstName: "Ankit",
	}
	fmt.Println(u2)
}
