package main

import (
	"fmt"
	"time"
)

func main() {
}

type Transaction interface {
	Apply(float64) float64
}

type Income struct {
	When        time.Time
	Amount      float64
	Description string
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
	When     time.Time
	Amount   float64
	Category string
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
