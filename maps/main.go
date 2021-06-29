package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   4.88,
		"pie":    34,
		"tomato": 3.4,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
}
