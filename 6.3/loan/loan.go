package loan

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrOutsideOfPeriod = errors.New("outside of acceptance period")
)

type InvalidAgeError struct {
	Age    int
	MinAge int
	MaxAge int
}

func (e *InvalidAgeError) Error() string {
	return fmt.Sprintf(
		"%d is outside of loan acceptable age of %d to %d",
		e.Age,
		e.MinAge,
		e.MaxAge,
	)
}

func ApplyForStudent(age int) (bool, error) {
	if age < 18 || age > 24 {
		return false, &InvalidAgeError{Age: age, MinAge: 18, MaxAge: 24}
	}

	month := time.Now().Month()
	if month < 6 || month > 8 {
		return false, ErrOutsideOfPeriod
	}

	return true, nil
}
