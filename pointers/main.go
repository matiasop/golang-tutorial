package main

import "fmt"

func updateName(x *string) {
	*x = "pedrito"
}

func main() {
	name := "juanin"

	// pointer
	fmt.Println(&name)
	// dereference a pointer
	fmt.Println(*&name)

	updateName(&name)

	fmt.Println(name)

}
