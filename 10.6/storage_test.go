package main

import (
	"bytes"
	"testing"
	"time"
)

func TestLoadAndSave(t *testing.T) {
	trx = append(trx, Income{When: time.Now(), Amount: 100.0, Description: "test"})
	trx = append(trx, Expense{When: time.Now(), Amount: 40.0, Category: "unittest"})

	var b []byte
	buf := bytes.NewBuffer(b)

	if err := save(buf, trx); err != nil {
		t.Fatal(err)
	}

	chk, err := load(buf)
	if err != nil {
		t.Fatal(err)
	} else if len(chk) != len(trx) {
		t.Errorf("expected chk len to be %d, got %d", len(trx), len(chk))
	}

	if i, ok := chk[0].(Income); !ok {
		t.Errorf("expected first trx to be an Income")
	} else if i.Amount != 100.0 {
		t.Errorf("expected 100.0 got %v", i.Amount)
	}

	if e, ok := chk[1].(Expense); !ok {
		t.Errorf("expected 2nd chk to be an Expense")
	} else if e.Amount != 40.0 {
		t.Errorf("expected 40.0, got %v", e.Amount)
	}
}
