package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := "4.54"
	p, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(p)
}
