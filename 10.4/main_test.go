package main

import (
	"testing"
	"time"
)

func TestCalculatingTransactions(t *testing.T) {
	allTrx := []Transaction{
		Income{When: time.Now(), Amount: 150.25, Description: "testing"},
		Expense{When: time.Now(), Amount: 30.0, Category: "online course"},
	}

	total := 0.0
	for _, trx := range allTrx {
		total = trx.Apply(total)
	}

	if total != 120.25 {
		t.Errorf("wanted 125.25, got %v", total)
	}
}
