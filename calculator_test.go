package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: -2, b: 2, want: -4},
		{a: 1, b: 0.25, want: 0.75},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: 1, want: -1},
		{a: 5, b: 10, want: 0.5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Divide(1,0): want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 9, want: 3},
		{a: 1, want: 1},
		{a: 4, want: 2},
		{2, 1.414214},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("Sqrt(%f): want no error for valid input, got %v", tc.a, err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Error("Sqrt(-1): want error for invalid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAddMultiple(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    []float64
		want float64
	}
	a := []float64{1, 2, 3, 4}
	b := []float64{3, 4, 5, 6}
	c := []float64{7, 8, 9, 8}
	testCases := []testCase{
		{a: a, want: 10},
		{a: b, want: 18},
		{a: c, want: 32},
	}
	for _, tc := range testCases {
		got, err := calculator.AddMultiple(tc.a...)
		if err != nil {
			t.Fatalf("AddMultiple(%f): want no error for valid input, got %v", tc.a, err)
		}
		if tc.want != got {
			t.Errorf("AddMultiple(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}

func TestDivideMultiple(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    []float64
		want float64
	}
	a := []float64{12, 4, 3}
	b := []float64{9, 3}
	c := []float64{8}
	testCases := []testCase{
		{a: a, want: 1},
		{a: b, want: 3},
		{a: c, want: 8},
	}
	for _, tc := range testCases {
		got, err := calculator.DivideMultiple(tc.a...)
		if err != nil {
			t.Fatalf("DivideMultiple(%f): want no error for valid input, got %v", tc.a, err)
		}
		if tc.want != got {
			t.Errorf("DivideMultiple(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}

func TestMultiplyMultiple(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    []float64
		want float64
	}
	a := []float64{1, 2, 3, 4}
	b := []float64{3, 4, 5, 6}
	c := []float64{7, 8, 9, 8}
	testCases := []testCase{
		{a: a, want: 24},
		{a: b, want: 360},
		{a: c, want: 4032},
	}
	for _, tc := range testCases {
		got, err := calculator.MultiplyMultiple(tc.a...)
		if err != nil {
			t.Fatalf("MultiplyMultiple(%f): want no error for valid input, got %v", tc.a, err)
		}
		if tc.want != got {
			t.Errorf("MultiplyMultiple(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}

func TestDivideMultipleInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.DivideMultiple(1, 2, 0, 9)
	if err == nil {
		t.Error("Divide(1,0): want error for invalid input, got nil")
	}
}
