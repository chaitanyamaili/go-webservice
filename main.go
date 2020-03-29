package main

func main() {
	// Contional loop.
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 2 {
			continue
		}
		println("Continue")
		if i == 3 {
			break
		}
	}

	// Loop to contion post clause
	for j := 0; j < 5; j++ {
		println(j)
	}

	// Infinite loop.
	var k int
	for {
		if k == 7 {
			break
		}
		println(k)
		k++
	}

	// Loops for the collections.
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
	wellKnownPort := map[string]int{"http": 80, "https": 443}
	for i, v := range wellKnownPort {
		println(i, v)
	}
	// Write only variable, if we need only value.
	for _, v := range wellKnownPort {
		println(v)
	}

	// Panic statement.
	println("Starting web server")
	panic("something bad happened")
	println("Web server started")
}
