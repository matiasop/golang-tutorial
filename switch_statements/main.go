package main

import "fmt"

func main() {
	opt := "s"
	switch opt {
	case "a":
		fmt.Println("you chose a")
	case "b":
		fmt.Println("you chose b")
	case "c":
		fmt.Println("you chose c")
	default:
		fmt.Println("not valid option")
	}
}
