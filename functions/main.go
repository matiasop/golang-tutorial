package main

import (
	"fmt"
	"strings"
)

func sayGreeting(n string, f func(string)) {
	fmt.Printf("good morning %v\n", n)
	f("hola")
}

func bye(n string) {
	fmt.Printf("%v bye", n)
}

func computation(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func getInitials(n string) (string, string) {
	// return multiple values
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	sayGreeting("juanin", bye)
	fmt.Println(computation([]int{2, 4, 6, 7, 32, 23}))
	a, b := getInitials("hola mundo")
	fmt.Println(a, b)
}
