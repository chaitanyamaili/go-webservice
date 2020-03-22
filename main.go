package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr2[:]
	fmt.Println(arr2, slice)

	// slice without defined length.
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	slice2 = append(slice2, 42, 47)
	slice2 = append(slice2, 12, 45)
	fmt.Println(slice2)

	// Working with maps.
	m := map[string]int{"foo": 12, "bar": 34, "a": 0}
	fmt.Println(m)
	// Update the value in map.
	m["foo"] = 123
	fmt.Println(m)
	// Delete the key from the map.
	delete(m, "foo")
	fmt.Println(m)

	// Working with struct data type.
	type user struct {
		ID        int
		firstname string
		lastname  string
	}
	var u user
	u.ID = 1
	u.firstname = "First"
	u.lastname = "Last"
	fmt.Println(u)

	u2 := user{ID: 1,
		firstname: "ram",
		lastname:  "govind",
	}
	fmt.Println(u2)
}
