package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Go inside webservice!")
	err := startWebServer(3000, 2)
	fmt.Println(err)

	// multiple returns.
	port, err := multpileReturns(200)
	fmt.Println(port, err)

	// ignore one of the return value.
	_, e := multpileReturns(200)
	fmt.Println(e)
}

func startWebServer(port, numberOfRetrives int) error {
	fmt.Println("Starting server...")
	// do importsnt things.
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retrives", numberOfRetrives)
	return errors.New("Something went wrong!")
}

func multpileReturns(port int) (int, error) {
	fmt.Println("port is", port)
	return port, nil
}
