package main

import "fmt"

func main() {
	// arrays
	var ages [3]int = [3]int{20, 25, 30}
	// var agesTwo = [3]int{20, 25, 30}
	ages[1] = 6

	fmt.Println(ages, len(ages))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 45
	scores = append(scores, 8)

	fmt.Println(scores)

	// slice ranges
	rangeOne := scores[1:3]
	fmt.Println(rangeOne)
}
