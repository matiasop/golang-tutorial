package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends"

	// return true if it contains the sub-string
	fmt.Println(strings.Contains(greeting, "there"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	// returns the position where it finds the sub-string
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{34, 234, 432, 643, 34}
	sort.Ints(ages)
	fmt.Println(ages)
}
