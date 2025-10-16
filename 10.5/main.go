package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var trx []Transaction

func main() {
	var file io.ReadWriteCloser
	var err error

	if fileExists("trx.dat") {
		file, err = os.Open("trx.dat")
		if err != nil {
			fmt.Println("unable to open file trx.dat: ", err)
			os.Exit(1)
		}

		trx, err = load(file)
		if err != nil {
			fmt.Print("unable to read trx.dat file: ", err)
			os.Exit(1)
		}
	} else {
		file, err = os.Create("trx.dat")
		if err != nil {
			fmt.Println("unable to create file trx.dat: ", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	// perform CLI action

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
		"%s\t+%f.2 - %s",
		i.When.Format(("2006/01/02 15:04")),
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
		"%s\t(%f.2) - %s",
		e.When.Format("2006/01/02 15:04"),
		e.Amount,
		e.Category,
	)
}
