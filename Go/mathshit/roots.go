package mathshit

import (
	"errors"
)

// first attempt at a Sqrt(int) function
func Sqrt(num int) (float64, error) {

	// check for negative input
	if num < 0 {
		return 0, errors.New("negative input")
	}

	if num == 0 {
		return 0, nil
	}

	// start at highest base
	base := 1.0
	for base < float64(num)/10 {
		base *= 10
	}

	var out float64

	for base > 0.00000001 {

		for out*out < float64(num) {
			out += base
		}

		out -= base

		base /= 10
	}

	out += base
	out = (float64(((int)(100000000 * out)))) / 100000000

	return out, nil

}
