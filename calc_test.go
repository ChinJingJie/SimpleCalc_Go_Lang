package main

import (
	"fmt"
	"testing"
)

func TestIsValidNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		wantErr  bool
	}{
		{"10", 10.0, false},         // Valid positive number
		{"-5.5", -5.5, false},       // Valid negative number
		{"3.14159", 3.14159, false}, // Valid float number
		{"", 0, true},               // Empty input (error expected)
		{"abc", 0, true},            // Invalid input type (error expected)
		{"*", 0, true},              // Invalid input type (error expected)
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result, err := isValidNumber(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error: %v, got error: %v", tc.wantErr, err)
				return
			}

			if !tc.wantErr && result != tc.expected {
				t.Errorf("Expected result: %g, got: %g", tc.expected, result)
			}
		})
	}
}

func TestIsValidSymbol(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"+", true},    // Valid add operator
		{"-", true},    // Valid minus operator
		{"*", true},    // Valid multiply operator
		{"/", true},    // Valid divide operator
		{"=", false},   // Invalid equal operator (error expected)
		{"", false},    // Empty input (error expected)
		{"add", false}, // Invalid input (error expected)
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := isValidSymbol(tc.input)

			if result != tc.expected {
				t.Errorf("Expected result: %t, got: %t", tc.expected, result)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		number1  float64
		number2  float64
		operator string
		expected float64
	}{
		{1, 27, "+", 28},  // Valid add equation
		{7, 9, "+", 16},   // Valid add equation
		{18, 13, "+", 31}, // Valid add equation
		{13, 42, "+", 55}, // Valid add equation
	}

	var calc Calculator

	for _, tc := range tests {
		calc.number1 = tc.number1
		calc.number2 = tc.number2
		calc.operator = tc.operator

		t.Run(fmt.Sprintf("Equation is `%g %s %g =`", calc.number1, calc.operator, calc.number2), func(t *testing.T) {
			result := calc.add()

			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	tests := []struct {
		number1  float64
		number2  float64
		operator string
		expected float64
	}{
		{1, 27, "-", -26},  // Valid minus equation
		{17, 9, "-", 8},    // Valid minus equation
		{18, 13, "-", 5},   // Valid minus equation
		{113, 42, "-", 71}, // Valid minus equation
	}

	var calc Calculator

	for _, tc := range tests {
		calc.number1 = tc.number1
		calc.number2 = tc.number2
		calc.operator = tc.operator

		t.Run(fmt.Sprintf("Equation is `%g %s %g =`", calc.number1, calc.operator, calc.number2), func(t *testing.T) {
			result := calc.minus()

			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		number1  float64
		number2  float64
		operator string
		expected float64
	}{
		{1, 27, "*", 27},   // Valid multiply equation
		{17, 0, "*", 0},    // Valid multiply equation
		{18, 13, "*", 234}, // Valid multiply equation
		{13, -2, "*", -26}, // Valid multiply equation
	}

	var calc Calculator

	for _, tc := range tests {
		calc.number1 = tc.number1
		calc.number2 = tc.number2
		calc.operator = tc.operator

		t.Run(fmt.Sprintf("Equation is `%g %s %g =`", calc.number1, calc.operator, calc.number2), func(t *testing.T) {
			result := calc.multiply()

			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		number1  float64
		number2  float64
		operator string
		expected float64
		wantErr  bool
	}{
		{144, 12, "/", 12, false}, // Valid divide equation
		{17, 17, "/", 1, false},   // Valid divide equation
		{32, 4, "/", 8, false},    // Valid divide equation
		{-25, 5, "/", -5, false},  // Valid divide equation
		{13, 0, "/", 0, true},     // Invalid divide equation
	}

	var calc Calculator

	for _, tc := range tests {
		calc.number1 = tc.number1
		calc.number2 = tc.number2
		calc.operator = tc.operator

		t.Run(fmt.Sprintf("Equation is `%g %s %g =`", calc.number1, calc.operator, calc.number2), func(t *testing.T) {
			result, err := calc.divide()

			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error: %v, got error: %v", tc.wantErr, err)
				return
			}

			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}
