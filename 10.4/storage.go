package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Ledger struct {
	In  []Income  `json:"in"`
	Out []Expense `json:"out"`
}

func load(filename string) ([]Transaction, error) {
	if !fileExists(filename) {
		return nil, nil
	}

	b, err := os.ReadFile(filename)
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

func save(filename string, trx []Transaction) error {
	var ledger Ledger
	for _, t := range trx {
		switch v := t.(type) {
		case Income:
			ledger.In = append(ledger.In, v)
		case Expense:
			ledger.Out = append(ledger.Out, v)
		}
	}

	if !fileExists(filename) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		return json.NewEncoder(file).Encode(ledger)
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(ledger)
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
