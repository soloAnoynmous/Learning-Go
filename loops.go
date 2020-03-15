package main

func main() {

	// loop till
	println("i..............")
	var i int
	for i<5 {
		println(i)
		i++
	}

	// loop till with post clause
	println("j..............")
	for j:=0;j<4;j++ {
		println(j)
	}

	// loop over collections
	println("collection_looping_1..............")
	slice := []int{1,2,3}
	for i:=0;i<len(slice);i++ {
		println(slice[i])
	}

	println("collection_looping_2..............")
	slices := []int{1,2,3}
	for i := range slices {
		println(i)
	}
}