package thefarm

import "fmt"

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	fodder, err := fc.FodderAmount(numCows)
    if err != nil {
        return 0.0, err
    }
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return (fodder*factor)/float64(numCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
    if numCows >0 {
        return DivideFood(fc, numCows)
    } else {
        return 0.0, &CowsError{message: "invalid number of cows"}
    }
}

func ValidateNumberOfCows(numCows int) (error) {
	switch {
        case numCows < 0: 
        	return &InvalidCowsError{num: numCows, message: "there are no negative cows"}
        case numCows == 0: 
        	return &InvalidCowsError{num: numCows, message: "no cows don't need food"}
    	default: 
        	return nil
    }
}

type CowsError struct {
    message string
}

func (e *CowsError) Error() string {
    return fmt.Sprintf("%s", e.message)
}
 
type InvalidCowsError struct {
    num int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %s", e.num, e.message)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
