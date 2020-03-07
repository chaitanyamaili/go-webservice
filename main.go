package main

import "fmt"

const (
	first = iota
	second
)

const (
	third = iota
	fourth
)

func main() {
	fmt.Println("Hello Go inside webservice!")

	// integer type
	var i int
	i = 12
	fmt.Println(i)

	// float32 type
	var f float32 = 3.14
	fmt.Println(f)

	// string type
	firstName := "Chaitanya"
	fmt.Println(firstName)

	// boolean type
	b := true
	fmt.Println(b)

	// complex type
	c := complex(3, 4)
	fmt.Println(c)

	// real and imagenary type
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// pointer type
	var fname *string = new(string)
	*fname = "Ram"
	fmt.Println(fname, *fname)

	ptr := &firstName
	fmt.Println(ptr, *ptr)
	firstName = "Ram2"
	fmt.Println(ptr, *ptr)

	// Constants
	const pi = 3.14
	fmt.Println(pi)

	// Constant expressions
	fmt.Println(first, second, third, fourth)
}
