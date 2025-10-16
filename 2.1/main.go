package main

import (
	"demo/calculator"
	"fmt"
)

func main() {
	res := calculator.Add(1, 2)
	fmt.Printf("result is %d\n", res)
}
