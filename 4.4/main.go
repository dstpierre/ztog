package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 4, 3, 2, 1}

	sort.Ints(s)
	fmt.Println(s)
}
