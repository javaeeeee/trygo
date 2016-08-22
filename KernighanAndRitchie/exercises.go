/*
Package KernighanAndRitchie provides implementation of some exercises from Kernighan & Ritchie book.
*/
package KernighanAndRitchie

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
)

// ErrorStepZero is the error message when step is 0
const ErrorStepZero = "ConvertFahrenheitToCelcius: step should be a positive number"

//ErrorLowBiggerThanHigh is the error message when lower limit is greater than higher
const ErrorLowBiggerThanHigh = "ConvertFahrenheitToCelcius: high should be greater than low"

// ErrorNotMultiple is the error message when the range is not multiple of the step
const ErrorNotMultiple = "ConvertFahrenheitToCelcius: difference between high and low should be multiple of the step"

/*
ConvertFahrenheitToCelcius converts temperature in Fahrenheit to Celcius.
*/
func ConvertFahrenheitToCelcius(low, high, step int) ([]float64, error) {
	// Check parameters validity
	if step <= 0 {
		return nil, errors.New(ErrorStepZero)
	} else if high <= low {
		return nil, errors.New(ErrorLowBiggerThanHigh)
	} else if (high-low)%step != 0 {
		return nil, errors.New(ErrorNotMultiple)
	}
	// calculate number of points in the result
	numPoints := (high-low)/step + 1
	// allocate an array; non-constant size is forbidden
	result := make([]float64, numPoints)
	// array index
	index := 0

	for fahr := low; fahr <= high; fahr += step {
		// use type conversion
		result[index] = 5. * float64(fahr-32.) / 9.
		index++
	}

	return result, nil
}

/*
Power computes m-th power of number n
*/
func Power(n float64, m int) float64 {
	if m == 0 {
		return 1.
	}
	result := float64(n)
	numIterations := int(math.Abs(float64(m)))
	for i := 1; i < numIterations; i++ {
		result *= float64(n)
	}
	if m < 0 {
		result = 1. / result
	}
	return result
}

/*
Power computes m-th power of number n using recursion
*/
func PowerRecursive(n float64, m int) float64 {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	if m == -1 {
		return 1. / n
	}
	if m > 0 {
		return n * PowerRecursive(n, m-1)
	}
	return PowerRecursive(n, m+1) / n
}

/*
CopyTextFile copies text file specified by file name src to a new file with name provided in dest parameter
*/
func CopyTextFile(src, dest string) error {
	if src == dest {
		return nil
	}
	data, err := ioutil.ReadFile(src)
	if err == nil {
		err = ioutil.WriteFile(dest, data, os.ModePerm)
	}

	// If error in either case is not nil, return the error
	if err != nil {
		return err
	}
	return nil
}
