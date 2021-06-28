package main

import "fmt"

func main() {
	// Print (without new line)
	fmt.Print("hello, ")
	fmt.Print("world \n")

	// Println (print with new line)
	fmt.Println("hello, world")

	name := "juanin"
	fmt.Println("my name is", name, ".")

	// Printf (formatted string), this needs a format specifier
	fmt.Printf("my age is %v.\n", name)
	fmt.Printf("my age is %q.\n", name)

	// Sprintf (save formatted string)
	var variable = fmt.Sprintf("my age is %v", name)
	fmt.Println(variable)
}
