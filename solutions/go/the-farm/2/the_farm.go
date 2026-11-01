package thefarm

import (
    "fmt"
	"errors"
)

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func DivideFood(fc FodderCalculator, noOfCows int) (float64, error) {
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    amount, err := fc.FodderAmount(noOfCows)
    if err != nil {
        return 0, err
    }
    return (amount * factor) / float64(noOfCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, noOfCows int) (float64, error) {
	if noOfCows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, noOfCows)
}

func ValidateNumberOfCows(noOfCows int) error {
	if noOfCows < 0 {
		return &InvalidCowsError{cows: noOfCows, message: "there are no negative cows"}
	}
	if noOfCows == 0 {
		return &InvalidCowsError{cows: noOfCows, message: "no cows don't need food"}
	}
    return nil
}

