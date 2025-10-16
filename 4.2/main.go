package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Dominic"] = 43
	m["Alice"] = 57

	v, ok := m["Bob"]
	if !ok {
		v = 100
		m["Bob"] = v
	}

	fmt.Println(v)
	fmt.Println(m)
}
