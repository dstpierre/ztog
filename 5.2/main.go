package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := div(2, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(res)
}

func sayHello(name string) {
	fmt.Println("hello", name)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
