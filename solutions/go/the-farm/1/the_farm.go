package thefarm

import (
    "fmt"
	"errors"
)

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
        msg := fmt.Sprintf("%d cows are invalid: there are no negative cows", noOfCows)
        return errors.New(msg)
    } else if noOfCows == 0 {
        return errors.New("0 cows are invalid: no cows don't need food")
    }
    return nil
}

