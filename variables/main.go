package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// strings
	var nameOne string = "manuel"
	var nameTwo = "jaime"
	var nameThree string // variable declaration
	nameFour := "Jorge"  // this shorthand cannot be used outside of a function

	nameThree = "juanin"
	fmt.Println(nameOne)
	fmt.Println(nameTwo)
	fmt.Println(nameThree)
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int32 = -4
	var numTwo uint64 = 5
	var numTwo float64 = 23.4321

	fmt.Println(numOne, numTwo, numTwo)

}
