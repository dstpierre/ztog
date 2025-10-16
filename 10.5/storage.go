package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Ledger struct {
	In  []Income  `json:"in"`
	Out []Expense `json:"out"`
}

func load(r io.Reader) ([]Transaction, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var ledger Ledger
	if err := json.Unmarshal(b, &ledger); err != nil {
		return nil, err
	}

	var trx []Transaction
	for _, t := range ledger.In {
		trx = append(trx, t)
	}

	for _, t := range ledger.Out {
		trx = append(trx, t)
	}

	return trx, nil
}

func save(w io.Writer, trx []Transaction) error {
	var ledger Ledger
	for _, t := range trx {
		switch v := t.(type) {
		case Income:
			ledger.In = append(ledger.In, v)
		case Expense:
			ledger.Out = append(ledger.Out, v)
		}
	}

	return json.NewEncoder(w).Encode(ledger)
}

// fileExists checks if a file exists and is not a directory.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil {
		return !info.IsDir()
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	fmt.Printf("Error accessing file %s: %v\n", filename, err)
	return false
}
