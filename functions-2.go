package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	isStarted := startWebServer(port, 2)
	fmt.Println(isStarted)

	err := startWebServerb(port, 2)
	fmt.Println(err)

	port, errs := startWebServerc(port)
	fmt.Println(port, errs)
}

// returning a value
func startWebServer(port int, numberOfRetries int) bool {
	fmt.Println("Starting server....")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return true
}

// returning an error
func startWebServerb(port int, numberOfRetries int) error {
	fmt.Println("Starting server....")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return errors.New("something went wrong")
}

// returning multiple values
func startWebServerc(port int) (int, error) {
	fmt.Println("Starting server....")
	fmt.Println("Server started on port", port)
	return port, errors.New("something went wrong")
}
