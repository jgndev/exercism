package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	NumberOfCows int
	Message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

// DivideFood takes a FoodCalculator and number of cows and returns the amount
// of food to feed each cow by determining the total amount of food and the
// fattening factor which are multiplied together and divided by the number of cows.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return (amount * factor) / float64(cows), nil
}

// ValidateInputAndDivideFood checks for invalid input before calling DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}

	err := ValidateNumberOfCows(cows)
	if err != nil {
		return 0.0, err
	}

	fpc, err := DivideFood(fc, cows)
	if err != nil {
		return 0.0, err
	}

	return fpc, nil
}

// ValidateNumberOfCows checks for an invalid number of cows.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			NumberOfCows: cows,
			Message:      "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			NumberOfCows: cows,
			Message:      "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
