package main

import "fmt"

func main() {
	age := 45
	if age < 30 {
		fmt.Println("hola")
	} else if age < 20 {
		fmt.Println("chao")
	} else {
		fmt.Println("otro caso")
	}
}
