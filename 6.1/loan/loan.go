package loan

import "errors"

func ApplyForStudent(age int) (bool, error) {
	if age < 18 || age > 24 {
		return false, errors.New("not between 18 and 24 years old")
	}

	return true, nil
}
