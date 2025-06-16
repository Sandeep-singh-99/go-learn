package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating maps

	m := make(map[string]string)

	m["name"] = "john"
	m["area"] = "backend"

	fmt.Println(m["name"])
	fmt.Println(m["area"])

	// IMP: if key does not exist in the map, it returns the zero value of the value type
	fmt.Println(m["phone"])

	n := make(map[string]int)
	n["age"] = 30
	fmt.Println(n["age"])

	fmt.Println(len(m)) // returns 0 since phone key does not exist

	delete(m, "area") // delete key from map
	fmt.Println(m)

	clear(m)

	fmt.Println(m) // map will be empty after clearing

	p := map[string]int{"a": 1, "b": 2, "c": 3} // shorthand for creating a map
	fmt.Println(p)

	k, ok := p["c"] // checking if key exists

	if ok {
		fmt.Println("Key 'c' exists with value:", k)
	} else {
		fmt.Println("Key 'c' does not exist")
	}

	m1 := map[string]string{"name": "Alice", "city": "Wonderland"} // another way to create a map
	m2 := map[string]string{"name": "Bob", "country": "Neverland"} // another map

	fmt.Println(maps.Equal(m1, m2)) // comparing two maps for equality
}
