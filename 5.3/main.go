package main

import "fmt"

func main() {
	fmt.Print(add(1, 2, 3, 4))

	me := "Dominic"
	me2 := doX(me)

	fmt.Print(me, me2)
}

func doX(s string) string {
	s += " changed"
	return s
}

func add(numbers ...int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}

	return result
}
