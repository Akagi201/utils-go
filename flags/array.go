package flags

import (
	"fmt"
)

// Array represents string array flag variable.
type Array []string

// String returns the string representation of the array.
func (a *Array) String() string {
	return fmt.Sprintf("%v", *a)
}

// Set appends element to the array.
func (a *Array) Set(value string) error {
	*a = append(*a, value)
	return nil
}
