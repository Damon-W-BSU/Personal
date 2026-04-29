package errordemo

import (
	"errors"
	"fmt"
)

// declare errors
var (
	errLow  = errors.New("less than 5")
	errHigh = errors.New("greater than 5")
	errNeg  = errors.New("negative")
)

// returns a specific error if num != 0
func isOne(num int) error {
	switch {
	case num == -1:
		return errNeg
	case num < 1:
		return errLow
	case num > 1:
		return errHigh
	default:
		return nil
	}
}

// returns wrapped error for isFive
func check(num int) error {
	err := isOne(num)
	// if error present, returns wrapped error
	if err != nil {
		return fmt.Errorf("check() error: %w", err)
	}
	return nil
}

func Demo() {
	for i := -2; i <= 2; i++ {
		err := check(i)
		var out string
		// unwrap check() error
		switch errors.Unwrap(err) {
		case errLow:
			out = "detected errLow" // specific error
		case errHigh:
			out = "detected errHigh" // specific error
		case nil:
			out = "no nested error" // check() error unwraps to nil
		default:
			out = "detected error" // general error
		}

		// output result
		fmt.Printf("check(%d):\n\terr: %s\n\tout: %s\n", i, err, out)
	}
}
