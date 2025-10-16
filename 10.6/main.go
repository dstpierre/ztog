package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var trx []Transaction

func main() {
	var err error

	if fileExists("trx.dat") {
		file, err := os.Open("trx.dat")
		if err != nil {
			fmt.Println("unable to open file trx.dat: ", err)
			os.Exit(1)
		}
		defer file.Close()

		trx, err = load(file)
		if err != nil {
			fmt.Print("unable to read trx.dat file: ", err)
			os.Exit(1)
		}
	}

	process(os.Args)

	file, err := os.Create("trx.dat")
	if err != nil {
		fmt.Println("unable to create the trx.dat file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := save(file, trx); err != nil {
		fmt.Println("error while saving trx.dat: ", err)
		os.Exit(1)
	}
}

type Transaction interface {
	Apply(float64) float64
}

type Income struct {
	When        time.Time `json:"when"`
	Amount      float64   `json:"amount"`
	Description string    `json:"desc"`
}

func (i Income) Apply(total float64) float64 {
	return total + i.Amount
}

func (i Income) String() string {
	return fmt.Sprintf(
		"%s\t+%.2f - %s",
		i.When.Format(("2006/01/02")),
		i.Amount,
		i.Description,
	)
}

type Expense struct {
	When     time.Time `json:"when"`
	Amount   float64   `json:"amount"`
	Category string    `json:"cat"`
}

func (e Expense) Apply(total float64) float64 {
	return total - e.Amount
}

func (e Expense) String() string {
	return fmt.Sprintf(
		"%s\t(%.2f) - %s",
		e.When.Format("2006/01/02"),
		e.Amount,
		e.Category,
	)
}

func process(args []string) {
	if len(args) == 3 {
		display(args[2])
		return
	} else if len(args) == 5 {
		processTrx(args[1], args[2], args[3], args[4])
		return
	}

	fmt.Println(`Usage:`)
	fmt.Println("  ztog [action] [params]")
	fmt.Println("")
	fmt.Println("\"\tztog income 2025-10-04 100.50 \"desc\"")
	fmt.Println("\tztog expense 2025-10-08 50 \"category\"")
	fmt.Println("")
	fmt.Println("\tztog month 2025-10")
}

func processTrx(action, date, amount, text string) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("invalid date: ", err)
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("invalid amount: ", err)
		os.Exit(1)
	}

	if action == "income" {
		trx = append(trx, Income{When: d, Amount: a, Description: text})
		return
	}

	trx = append(trx, Expense{When: d, Amount: a, Category: text})
}

func display(date string) {
	d, err := time.Parse("2006-01", date)
	if err != nil {
		fmt.Println("invalid year-month: ", err)
		os.Exit(1)
	}

	for _, t := range filter(d) {
		fmt.Println(t)
	}
}

func filter(d time.Time) []Transaction {
	var filtered []Transaction
	for _, t := range trx {
		switch v := t.(type) {
		case Income:
			if v.When.Year() == d.Year() && v.When.Month() == d.Month() {
				filtered = append(filtered, t)
			}
		case Expense:
			if v.When.Year() == d.Year() && v.When.Month() == d.Month() {
				filtered = append(filtered, t)
			}
		}
	}

	return filtered
}
