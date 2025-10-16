package order

func ApplyDiscount(total float64) float64 {
	if total < 0 {
		return 0
	}

	if total >= 100 {
		return total - (total * 0.15)
	}

	if total > 50 {
		return total - (total * 0.1)
	}

	return total
}
