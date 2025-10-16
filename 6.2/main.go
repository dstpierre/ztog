package main

import (
	"errors"
	"flag"
	"fmt"
	"ztog/loan"
)

func main() {
	var age int
	flag.IntVar(&age, "age", 0, "Age")
	flag.Parse()

	if err := start(age); err != nil {
		if err == loan.ErrOutsideOfPeriod {
			fmt.Println("this does not match anymore.")
		} else if errors.Is(err, loan.ErrOutsideOfPeriod) {
			fmt.Println("this is the correct way to unwrap error")
		}
		fmt.Println(err)
		return
	}

	fmt.Println("you're approved!")
}

func start(age int) error {
	if _, err := loan.ApplyForStudent(age); err != nil {
		if err == loan.ErrOutsideOfPeriod {
			return fmt.Errorf("please apply between june and august: %w", err)
		}
		return fmt.Errorf("cannot apply for loan: %w", err)
	}

	return nil
}
