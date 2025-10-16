package order_test

import (
	"testing"
	"ztg/order"
)

type DiscountCases struct {
	Name   string
	Total  float64
	Result float64
}

func TestApplyDiscount(t *testing.T) {
	cases := []DiscountCases{
		{Name: "Order for > 100", Total: 110.0, Result: 93.5},
		{Name: "Order for > 50", Total: 65, Result: 58.5},
		{Name: "Order for < 50", Total: 21.5, Result: 21.5},
		{Name: "Negative number should return 0", Total: -1, Result: 0},
		{Name: "Order at 100 should get the 15% discount", Total: 100.0, Result: 85.0},
		{Name: "Order at 50 should not get any discount", Total: 50.0, Result: 50.0},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			if res := order.ApplyDiscount(tc.Total); tc.Result != res {
				t.Errorf("wanted %v, got %v", tc.Result, res)
			}
		})
	}
}
