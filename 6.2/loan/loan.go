package loan

import (
	"errors"
	"time"
)

var (
	ErrOutsideOfPeriod = errors.New("outside of acceptance period")
)

func ApplyForStudent(age int) (bool, error) {
	if age < 18 || age > 24 {
		return false, errors.New("not between 18 and 24 years old")
	}

	month := time.Now().Month()
	if month < 6 || month > 8 {
		return false, ErrOutsideOfPeriod
	}

	return true, nil
}
