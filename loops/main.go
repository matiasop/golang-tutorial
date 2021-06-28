package main

import "fmt"

func main() {
	x := 0

	// while loop
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"aaa", "bbb", "ccc", "ddd"}
	for index, value := range names {
		fmt.Println(index, value)
	}
}
