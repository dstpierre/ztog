package main

import (
	"flag"
	"fmt"
	"ztog/loan"
)

func main() {
	var age int
	flag.IntVar(&age, "age", 0, "Age")
	flag.Parse()

	if err := start(age); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("you're approved!")
}

func start(age int) error {
	if _, err := loan.ApplyForStudent(age); err != nil {
		return fmt.Errorf("cannot apply for loan: %w", err)
	}

	return nil
}
