package main

import (
	"fmt"
)

// CustomErrorCode represents a custom error code.
type CustomErrorCode int

// Error returns the error message for the custom error code.
func (e CustomErrorCode) Error() string {
	switch e {
	case 404:
		return "RESOURCE_NOT_FOUND"
	case 502:
		return "BAD_GATEWAY"
	// Add more cases as needed
	default:
		return fmt.Sprintf("Unknown custom error code: %d", e)
	}
}

// Function that might return the custom error code
func someOperation() error {
	// Perform some operation that may result in an error
	// For example, return a custom error code 1
	return CustomErrorCode(404)
}

func main() {
	if err := someOperation(); err != nil {
		fmt.Println("Error:", err)
	}
}
