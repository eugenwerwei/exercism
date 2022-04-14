package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf(e.message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			if amount < 0 {
				return 0, errors.New("negative fodder")
			}
			amount *= 2
		} else {
			return 0, err
		}
	} else if amount < 0 {
		return 0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
	}

	return amount / float64(cows), nil
}
