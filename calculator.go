// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a from b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and
// returns the result of multiplying a from b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero not allowed")
	}
	return a / b, nil
}

// Sqrt takes two numbers a and b, and
// returns the result of square root.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative square root not allowed")
	}
	return math.Sqrt(a), nil
}

// AddMultiple takes a variable number
// of parameters and adds them

func AddMultiple(a ...float64) (float64, error) {
	if len(a) == 0 {
		return 0, nil
	}
	multipleAdd := a[0]
	for _, numb := range a[1:] {
		multipleAdd += numb
	}
	return multipleAdd, nil
}

// DivideMultiple takes a variable number
// of parameters and adds them

func DivideMultiple(a ...float64) (float64, error) {
	if len(a) == 0 {
		return 0, nil
	}
	var multipleDivide float64
	multipleDivide = a[0]
	for _, numb := range a[1:] {
		if numb == 0 {
			return 0, errors.New("divide by zero not allowed")
		}
		multipleDivide /= numb
	}
	//for i, numb := range a {
	//	if i == 0 {
	//		multipleDivide = numb
	//	}
	//	if i > 0 {
	//		multipleDivide /= numb
	//	}
	//
	//}
	return multipleDivide, nil
}

// DivideMultiple takes a variable number
// of parameters and adds them

func MultiplyMultiple(a ...float64) (float64, error) {
	if len(a) == 0 {
		return 0, nil
	}
	multipleMultiply := a[0]
	for _, numb := range a[1:] {
		multipleMultiply *= numb
	}
	return multipleMultiply, nil
}
