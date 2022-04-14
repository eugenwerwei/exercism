package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError interface {
	// wip
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
	}

	return amount / float64(cows), nil
}
