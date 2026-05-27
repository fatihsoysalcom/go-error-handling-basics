package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrDivisionByZero is a custom error variable for division by zero.
// This demonstrates creating a specific error type for a known condition.
var ErrDivisionByZero = errors.New("cannot divide by zero")

// safeDivide takes two integers and returns their division result or an error.
// This function exemplifies Go's idiomatic approach of returning (value, error).
func safeDivide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		// If an error condition is met, return the zero value for the result type
		// and the specific error.
		return 0, ErrDivisionByZero
	}
	// If no error, return the result and nil for the error.
	return numerator / denominator, nil
}

// parseAndCalculate takes two string inputs, attempts to parse them as integers,
// and then performs a safe division. It demonstrates error propagation and wrapping.
func parseAndCalculate(numStr, denStr string) (int, error) {
	// Step 1: Parse numerator string to int
	numerator, err := strconv.Atoi(numStr)
	if err != nil {
		// If an error occurs, immediately check it and return.
		// fmt.Errorf with %w wraps the original error, adding context.
		return 0, fmt.Errorf("invalid numerator '%s': %w", numStr, err)
	}

	// Step 2: Parse denominator string to int
	denominator, err := strconv.Atoi(denStr)
	if err != nil {
		// Propagate the error, adding context.
		return 0, fmt.Errorf("invalid denominator '%s': %w", denStr, err)
	}

	// Step 3: Perform safe division
	result, err := safeDivide(numerator, denominator)
	if err != nil {
		// Propagate the error from safeDivide, adding context.
		return 0, fmt.Errorf("division failed: %w", err)
	}

	// All operations successful, return result and nil error.
	return result, nil
}

func main() {
	// Test Case 1: Valid input
	fmt.Println("--- Test Case 1: Valid Input ---")
	result, err := parseAndCalculate("10", "2")
	if err != nil {
		// This block is skipped for successful operations.
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of 10 / 2: %d\n", result)
	}
	fmt.Println()

	// Test Case 2: Invalid numerator (parsing error)
	fmt.Println("--- Test Case 2: Invalid Numerator ---")
	result, err = parseAndCalculate("abc", "2")
	if err != nil {
		// The error is caught and printed.
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of abc / 2: %d\n", result)
	}
	fmt.Println()

	// Test Case 3: Invalid denominator (parsing error)
	fmt.Println("--- Test Case 3: Invalid Denominator ---")
	result, err = parseAndCalculate("10", "xyz")
	if err != nil {
		// The error is caught and printed.
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of 10 / xyz: %d\n", result)
	}
	fmt.Println()

	// Test Case 4: Division by zero (custom error)
	fmt.Println("--- Test Case 4: Division by Zero ---")
	result, err = parseAndCalculate("10", "0")
	if err != nil {
		// errors.Is is used to check if the error chain contains a specific error type.
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Printf("Specific Error: %v (Division by zero detected)\n", err)
		} else {
			fmt.Printf("General Error: %v\n", err)
		}
	} else {
		fmt.Printf("Result of 10 / 0: %d\n", result)
	}
	fmt.Println()

	// Test Case 5: Another valid input
	fmt.Println("--- Test Case 5: Another Valid Input ---")
	result, err = parseAndCalculate("100", "5")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of 100 / 5: %d\n", result)
	}
	fmt.Println()
}
