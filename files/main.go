package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("I'm saving this string")
	// filename, data, file permissions
	err := os.WriteFile("./filename.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("saved")
}
