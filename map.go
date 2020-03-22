package main

import "fmt"

func map_tour() {
	m := make(map[string]int)
	m["abc"] = 123
	m["xyz"] = 456

	fmt.Println("len(m)", len(m))

	for k, v := range(m) {
		fmt.Println("map", k, "=>", v)
	}
}

func main() {
	fmt.Println("Hello, go")

	map_tour()
}
