package jwt

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
)

func (r *repository) findNumbers(numbers []int64) (int64, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	if numbers[len(numbers)-1] == int64(len(numbers)-1) {
		return int64(len(numbers)), nil
	}
	for i, n := range numbers {
		if n != int64(i) {
			return int64(i), nil
		}
	}

	return 0, errors.ErrAuthNumberAssignmentFailedRaw
}
